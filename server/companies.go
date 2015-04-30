package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/gooby"
)

type CompanyResource struct {
	Companies map[string]gooby.Company
}

func RegisterCompanies(c *restful.Container) *CompanyResource {
	h := &CompanyResource{
		Companies: map[string]gooby.Company{
			"Bloodhound Gang": gooby.Company{Name: "Bloodhound Gang"},
		},
	}

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
	companies := make([]gooby.Company, len(h.Companies))
	i := 0
	for _, c := range h.Companies {
		companies[i] = c
		i += 1
	}
	res.WriteEntity(companies)
}

func (h *CompanyResource) create(req *restful.Request, res *restful.Response) {
	c := new(gooby.Company)
	req.ReadEntity(c)

	h.Companies[c.Name] = *c
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(c)
}

func (h *CompanyResource) del(req *restful.Request, res *restful.Response) {
	if _, ok := h.Companies[req.PathParameter("name")]; !ok {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	} else {
		delete(h.Companies, req.PathParameter("name"))
		res.WriteHeader(http.StatusNoContent)
	}
}

func (h *CompanyResource) get(req *restful.Request, res *restful.Response) {
	if c, ok := h.Companies[req.PathParameter("name")]; ok {
		res.WriteEntity(c)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}
