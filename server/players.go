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

	ws.Filter(auth.AuthorizationFilter)

	ws.Path("/players").
		Doc("Manage game players").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.PUT("/").To(h.save).
		Doc("Save a player").
		Operation("save").
		Reads(game.Player{}))

	ws.Route(ws.GET("/{name}").To(h.get).
		Doc("Get a player by name").
		Operation("get").
		Param(ws.PathParameter("name", "the player's name").DataType("string")).
		Writes(game.Player{}))

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
