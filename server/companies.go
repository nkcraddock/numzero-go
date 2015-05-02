package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/gooby"
)

type CompanyResource struct {
	store *gooby.Store
}

func RegisterCompanies(c *restful.Container) *CompanyResource {
	h := &CompanyResource{store: gooby.NewStore("Bloodhound Gang")}

	ws := new(restful.WebService)

	ws.Path("/companies").
		Doc("Manage companies").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/").To(h.all).
		Doc("get all companies").
		Operation("all").
		Writes([]gooby.Company{}))

	ws.Route(ws.GET("/{name}").To(h.get).
		Doc("get a single company").
		Operation("get").
		Param(ws.PathParameter("name", "name of the company").DataType("string")).
		Writes(gooby.Company{}))

	ws.Route(ws.DELETE("/{name}").To(h.del).
		Doc("delete a single company").
		Operation("delete").
		Param(ws.PathParameter("name", "name of the company").DataType("string")))

	ws.Route(ws.POST("/").To(h.create).
		Doc("create a new company").
		Operation("create").
		Reads(gooby.Company{}))

	c.Add(ws)

	return h
}

func (h *CompanyResource) all(req *restful.Request, res *restful.Response) {
	companies := h.store.GetCompanies()
	res.WriteEntity(companies)
}

func (h *CompanyResource) create(req *restful.Request, res *restful.Response) {
	c := new(gooby.Company)
	req.ReadEntity(c)

	h.store.SaveCompany(c)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(c)
}

func (h *CompanyResource) del(req *restful.Request, res *restful.Response) {
	if ok := h.store.DeleteCompany(req.PathParameter("name")); !ok {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		res.WriteHeader(http.StatusNoContent)
	}
}

func (h *CompanyResource) get(req *restful.Request, res *restful.Response) {
	if c, ok := h.store.GetCompany(req.PathParameter("name")); ok {
		res.WriteEntity(c)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
