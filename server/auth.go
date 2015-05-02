package server

import (
	"net/http"
	"strings"

	"code.google.com/p/go-uuid/uuid"

	"github.com/emicklei/go-restful"
)

type AuthResource struct {
	sessions map[string]TokenResponse
}

type TokenRequest struct {
	GrantType string `json:"grant_type"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	ClientId  string `json:"client_id"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	IdToken     string `json:"id_token"`
}

func RegisterAuth(c *restful.Container) *AuthResource {
	h := &AuthResource{sessions: make(map[string]TokenResponse)}
	c.Filter(h.AuthorizationFilter)

	ws := new(restful.WebService)

	ws.Path("/auth").
		Doc("Manages authorization").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.POST("/token").To(h.createToken).
		Doc("create a new access token").
		Operation("createToken").
		Reads(TokenRequest{}).
		Writes(""))

	c.Add(ws)

	return h
}

func (h *AuthResource) AuthorizationFilter(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {
	// auth/token is exempt
	if req.SelectedRoutePath() == "/auth/token" {
		chain.ProcessFilter(req, res)
		return
	}

	token := req.Request.Header.Get("Authorization")

	if strings.HasPrefix(token, "Bearer ") {
		if _, ok := h.sessions[token[7:]]; ok {
			chain.ProcessFilter(req, res)
			return
		}
	}

	res.AddHeader("WWW-Authenticate", `OAuth realm="http://localhost:3001/"`)
	res.WriteErrorString(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *AuthResource) createToken(req *restful.Request, res *restful.Response) {
	tr := new(TokenRequest)
	req.ReadEntity(tr)

	if tr.Username == "username" && tr.Password == "password" {
		token := &TokenResponse{
			AccessToken: "1234",
			TokenType:   "Bearer",
			ExpiresIn:   3600,
			IdToken:     uuid.New(),
		}

		h.sessions[token.IdToken] = *token

		res.WriteHeader(http.StatusFound)
		res.WriteEntity(token)
	} else {
		res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
}
