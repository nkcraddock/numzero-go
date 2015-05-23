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
	event := new(game.Event)
	req.ReadEntity(event)

	h.store.SaveEvent(event)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(event)
}

func (h *EventsResource) get(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	if event, err := h.store.GetEvent(id); err == nil {
		res.WriteEntity(event)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
