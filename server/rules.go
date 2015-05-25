package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero/game"
)

type RulesResource struct {
	store game.Store
}

func RegisterRulesResource(rootPath string, c *restful.Container, store game.Store, auth *AuthResource) *RulesResource {
	h := &RulesResource{store: store}

	ws := new(restful.WebService)

	ws.Path(rootPath + "/rules").
		Doc("Manage game rules").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.PUT("/").To(h.save).
		Doc("Save a rule").
		Operation("save").
		Reads(game.Rule{}))

	ws.Route(ws.GET("/").To(h.list).
		Doc("List all rules").
		Operation("list").
		Writes([]game.Rule{}))

	ws.Route(ws.GET("/{code}").To(h.get).
		Doc("Get a rule").
		Operation("get").
		Param(ws.PathParameter("code", "the rule's code").DataType("string")).
		Writes(game.Rule{}))

	c.Add(ws)

	return h
}

func (h *RulesResource) save(req *restful.Request, res *restful.Response) {
	if err := h.store.Open(); err != nil {
		if handleError(err, "", http.StatusBadRequest, res) {
			return
		}
	}
	defer h.store.Close()

	rule := new(game.Rule)
	req.ReadEntity(rule)

	h.store.SaveRule(rule)
	res.WriteHeader(http.StatusCreated)
	res.WriteEntity(rule)
}

func (h *RulesResource) get(req *restful.Request, res *restful.Response) {
	if err := h.store.Open(); err != nil {
		if handleError(err, "", http.StatusBadRequest, res) {
			return
		}
	}
	defer h.store.Close()

	code := req.PathParameter("code")
	if rule, err := h.store.GetRule(code); err == nil {
		res.WriteEntity(rule)
	} else {
		res.WriteErrorString(http.StatusNotFound, http.StatusText(http.StatusNotFound))
	}
}

func (h *RulesResource) list(req *restful.Request, res *restful.Response) {
	if err := h.store.Open(); err != nil {
		if handleError(err, "", http.StatusBadRequest, res) {
			return
		}
	}
	defer h.store.Close()

	rules, _ := h.store.ListRules()
	res.WriteEntity(rules)
}
