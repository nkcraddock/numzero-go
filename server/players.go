package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero/game"
)

type PlayersResource struct {
	store game.Store
}

func RegisterPlayersResource(c *restful.Container, store game.Store, auth *AuthResource) *PlayersResource {
	h := &PlayersResource{store: store}

	ws := new(restful.WebService)

	ws.Path("/players").
		Doc("Manage game players").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.PUT("/").To(h.save).
		Doc("Save a player").
		Operation("save").
		Reads(game.Player{}))

	ws.Route(ws.GET("/").To(h.list).
		Doc("Get a list of players").
		Operation("list").
		Writes([]game.Player{}))

	ws.Route(ws.GET("/{name}").To(h.get).
		Doc("Get a player by name").
		Operation("get").
		Param(ws.PathParameter("name", "the player's name").DataType("string")).
		Writes(game.Player{}))

	ws.Route(ws.POST("/{name}/activities").To(h.newActivity).
		Doc("Add activity for a player").
		Operation("newActivity").
		Param(ws.PathParameter("name", "the player's name").DataType("string")).
		Reads(Activity{}))

	ws.Route(ws.GET("/{name}/activities").To(h.listActivities).
		Doc("List activities for a player").
		Operation("listActivities").
		Param(ws.PathParameter("name", "the player's name").DataType("string")).
		Writes([]Activity{}))

	c.Add(ws)

	return h
}

func (h *PlayersResource) save(req *restful.Request, res *restful.Response) {
	player := new(game.Player)
	req.ReadEntity(player)

	h.store.SavePlayer(player)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(player)
}

func (h *PlayersResource) get(req *restful.Request, res *restful.Response) {
	name := req.PathParameter("name")
	if player, err := h.store.GetPlayer(name); err == nil {
		res.WriteEntity(player)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}

func (h *PlayersResource) list(req *restful.Request, res *restful.Response) {
	p, err := h.store.ListPlayers()
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
	res.WriteEntity(p)
}

func (h *PlayersResource) listActivities(req *restful.Request, res *restful.Response) {
	name := req.PathParameter("name")
	player, err := h.store.GetPlayer(name)
	if err != nil {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	acts := make([]Activity, len(player.Events))
	for i, evt := range player.Events {
		acts[i] = eventToAct(evt)
	}

	res.WriteEntity(acts)
}

func (h *PlayersResource) newActivity(req *restful.Request, res *restful.Response) {
	name := req.PathParameter("name")
	player, err := h.store.GetPlayer(name)
	if err != nil {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}

	act := new(Activity)
	req.ReadEntity(act)
	evt, err := h.actToEvent(act)

	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	}

	player.AddEvent(evt)
	h.store.SavePlayer(player)
	res.WriteHeader(http.StatusOK)
}

// eventToAct converts a game.Event to an Activity resource
func eventToAct(evt game.Event) Activity {
	act := Activity{
		Description: evt.Description,
		Url:         evt.Url,
		Scores:      make(map[string]int),
	}

	for _, s := range evt.Scores {
		act.Scores[s.Rule.Code] = s.Times
	}

	return act
}

// actToEvent converts an Activity resource to a game.Event
func (h *PlayersResource) actToEvent(act *Activity) (*game.Event, error) {
	scores := make([]game.Score, len(act.Scores))
	total := 0
	i := 0
	for code, cnt := range act.Scores {
		rule, err := h.store.GetRule(code)
		if err != nil {
			return nil, err
		}
		scores[i] = game.Score{&rule, cnt}
		total += rule.Points * cnt
		i += 1
	}
	evt := &game.Event{
		Description: act.Description,
		Url:         act.Url,
		Scores:      scores,
		Total:       total,
	}
	return evt, nil
}
