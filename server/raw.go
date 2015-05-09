package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
)

type RawResource struct {
	store numzero.Repo
}

func RegisterRaw(c *restful.Container, store numzero.Repo, auth *AuthResource) *RawResource {
	h := &RawResource{store: store}

	ws := new(restful.WebService)

	ws.Filter(auth.AuthorizationFilter)

	ws.Path("/raw").
		Doc("Manage raw").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/").To(h.all).
		Doc("get all raw").
		Operation("all").
		Writes([]numzero.RawBuildData{}))

	ws.Route(ws.GET("/{id}").To(h.get).
		Doc("get a single raw").
		Operation("get").
		Param(ws.PathParameter("id", "id of the raw").DataType("string")).
		Writes(numzero.RawBuildData{}))

	ws.Route(ws.DELETE("/{id}").To(h.del).
		Doc("delete a single raw").
		Operation("delete").
		Param(ws.PathParameter("id", "id of the raw").DataType("string")))

	ws.Route(ws.POST("/").To(h.create).
		Doc("create a new raw").
		Operation("create").
		Reads(numzero.RawBuildData{}))

	c.Add(ws)

	return h
}

func (h *RawResource) all(req *restful.Request, res *restful.Response) {
	raw := h.store.List()
	res.WriteEntity(raw)
}

func (h *RawResource) create(req *restful.Request, res *restful.Response) {
	c := new(numzero.RawBuildData)
	req.ReadEntity(c)

	h.store.Save(c)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(c)
}

func (h *RawResource) del(req *restful.Request, res *restful.Response) {
	if err := h.store.Delete(req.PathParameter("id")); err != nil {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

func (h *RawResource) get(req *restful.Request, res *restful.Response) {
	if c, err := h.store.Get(req.PathParameter("id")); err != nil {
		res.WriteEntity(c)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
