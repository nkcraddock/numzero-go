package server

import (
	"net/http"
	"strconv"

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
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

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

	ws.Route(ws.GET("/{name}/events").To(h.getEvents).
		Doc("Get a list of events for the player").
		Operation("getEvents").
		Param(ws.PathParameter("name", "the player's name").DataType("string")).
		Param(ws.QueryParameter("cnt", "the max number of events to return").DataType("string")).
		Writes([]game.Event{}))

	c.Add(ws)

	return h
}

func (h *PlayersResource) save(req *restful.Request, res *restful.Response) {
	h.store.Open()
	defer h.store.Close()

	player := new(game.Player)
	req.ReadEntity(player)

	h.store.SavePlayer(player)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(player)
}

func (h *PlayersResource) get(req *restful.Request, res *restful.Response) {
	h.store.Open()
	defer h.store.Close()

	h.store.Open()
	defer h.store.Close()
	name := req.PathParameter("name")
	if player, err := h.store.GetPlayer(name); err == nil {
		res.WriteEntity(player)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}

func (h *PlayersResource) list(req *restful.Request, res *restful.Response) {
	const layout = "01/02/2006 3:04pm"
	h.store.Open()
	defer h.store.Close()

	players, err := h.store.ListPlayers()
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	model := make([]map[string]interface{}, len(players))

	for i, p := range players {
		model[i] = map[string]interface{}{
			"name":  p.Name,
			"score": p.Score,
			"image": p.Image,
		}

		if events, err := h.store.GetPlayerEvents(p.Name, 1); err == nil && len(events) > 0 {
			evt := events[0]
			model[i]["lastEvent"] = map[string]interface{}{
				"id":     evt.Id,
				"desc":   evt.Description,
				"scores": evt.Scores,
				"date":   evt.Date,
				"total":  evt.Total,
			}
		}
	}

	res.WriteEntity(model)
}

func (h *PlayersResource) getEvents(req *restful.Request, res *restful.Response) {
	h.store.Open()
	defer h.store.Close()

	name := req.PathParameter("name")
	cnt, _ := strconv.Atoi(req.QueryParameter("cnt"))

	p, err := h.store.GetPlayerEvents(name, int64(cnt))
	if err != nil {
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
		return
	}

	res.WriteEntity(p)
}
