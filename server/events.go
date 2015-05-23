package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero/game"
)

type EventsResource struct {
	store game.Store
}

func RegisterEventsResource(c *restful.Container, store game.Store, auth *AuthResource) *EventsResource {
	h := &EventsResource{store: store}

	ws := new(restful.WebService)

	ws.Path("/events").
		Doc("Manage game events").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/").To(h.save).
		Doc("Save a event").
		Operation("save").
		Reads(game.Event{}))

	ws.Route(ws.GET("/{id}").To(h.get).
		Doc("Get an event by id").
		Operation("get").
		Param(ws.PathParameter("id", "the event's guid").DataType("string")).
		Writes(game.Event{}))

	c.Add(ws)

	return h
}

func (h *EventsResource) save(req *restful.Request, res *restful.Response) {
	if err := h.store.Open(); err != nil {
		if handleError(err, "", http.StatusBadRequest, res) {
			return
		}
	}
	defer h.store.Close()

	event := new(game.Event)
	req.ReadEntity(event)

	player, err := h.store.GetPlayer(event.Player)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, "Player not found.")
		return
	}

	event.Total = 0
	for _, score := range event.Scores {
		rule, err := h.store.GetRule(score.Rule)
		if err != nil {
			res.WriteErrorString(http.StatusBadRequest, "Invalid rule")
			return
		}
		event.Total += rule.Points * score.Times
	}

	player.AddEvent(event)
	h.store.SavePlayer(player)

	h.store.SaveEvent(event)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(event)
}

func (h *EventsResource) get(req *restful.Request, res *restful.Response) {
	err := h.store.Open()
	if handleError(err, "", http.StatusInternalServerError, res) {
		return
	}
	defer h.store.Close()

	id := req.PathParameter("id")
	if event, err := h.store.GetEvent(id); err == nil {
		res.WriteEntity(event)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
