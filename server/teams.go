package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
)

type TeamResource struct {
	store numzero.Store
}

func RegisterTeams(c *restful.Container, store numzero.Store, auth *AuthResource) *TeamResource {
	h := &TeamResource{store: store}

	ws := new(restful.WebService)

	ws.Filter(auth.AuthorizationFilter)

	ws.Path("/teams").
		Doc("Manage teams").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/").To(h.all).
		Doc("get all teams").
		Operation("all").
		Writes([]numzero.Team{}))

	ws.Route(ws.GET("/{name}").To(h.get).
		Doc("get a single team").
		Operation("get").
		Param(ws.PathParameter("name", "name of the team").DataType("string")).
		Writes(numzero.Team{}))

	ws.Route(ws.DELETE("/{name}").To(h.del).
		Doc("delete a single team").
		Operation("delete").
		Param(ws.PathParameter("name", "name of the team").DataType("string")))

	ws.Route(ws.POST("/").To(h.create).
		Doc("create a new team").
		Operation("create").
		Reads(numzero.Team{}))

	c.Add(ws)

	return h
}

func (h *TeamResource) all(req *restful.Request, res *restful.Response) {
	teams := h.store.GetTeams()
	res.WriteEntity(teams)
}

func (h *TeamResource) create(req *restful.Request, res *restful.Response) {
	c := new(numzero.Team)
	req.ReadEntity(c)

	h.store.SaveTeam(c)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(c)
}

func (h *TeamResource) del(req *restful.Request, res *restful.Response) {
	if ok := h.store.DeleteTeam(req.PathParameter("name")); !ok {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

func (h *TeamResource) get(req *restful.Request, res *restful.Response) {
	if c, ok := h.store.GetTeam(req.PathParameter("name")); ok {
		res.WriteEntity(c)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
