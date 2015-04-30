package server

import "github.com/emicklei/go-restful"

type CompanyResource struct {
}

func RegisterCompanies(c *restful.Container) *CompanyResource {
	h := &CompanyResource{}

	ws := new(restful.WebService)

	ws.Path("/companies").
		Doc("Manage companies").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/").To(h.all).
		Doc("get all companies").
		Operation("all").
		Writes([]string{}))

	c.Add(ws)

	return h
}

func (h *CompanyResource) all(req *restful.Request, res *restful.Response) {
	res.WriteEntity([]string{"chicken", "sandwich"})
}
