package server

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero/game"
)

type EventsResource struct {
	store game.Store
	gm    *game.GM
	hook  string
}

func RegisterEventsResource(c *restful.Container, store game.Store, auth *AuthResource, hook string) *EventsResource {
	h := &EventsResource{store: store, gm: game.NewGameMaster(store), hook: hook}

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

	result, err := h.gm.AddEvent(event)
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}

	if len(result.Achievements) > 0 && h.hook != "" {
		for _, a := range result.Achievements {
			txt := fmt.Sprintf("@%s earned an achievement: %s", result.Player, a.Name)

			_, err := request("POST", h.hook, map[string]interface{}{
				"text": txt,
			})

			if err != nil {
				res.WriteErrorString(http.StatusBadRequest, err.Error())
				return
			}

		}

	}

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

func request(verb, url string, thing interface{}) (*http.Response, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	pJson, _ := json.Marshal(thing)
	req, _ := http.NewRequest(verb, url, bytes.NewBuffer(pJson))
	req.Header.Set("content-type", "application/json")
	return client.Do(req)
}
