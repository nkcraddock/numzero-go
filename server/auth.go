package server

import (
	"net/http"
	"strings"

	"code.google.com/p/go-uuid/uuid"

	"github.com/dgrijalva/jwt-go"
	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/gooby"
)

type AuthResource struct {
	signingKey []byte
	store      *gooby.Store
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

func RegisterAuth(c *restful.Container, store *gooby.Store, signingKey []byte) *AuthResource {
	h := &AuthResource{store: store, signingKey: signingKey}
	c.Filter(h.AuthorizationFilter)

	ws := new(restful.WebService)

	ws.Path("/auth").
		Doc("Manages authorization").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.POST("/token").To(h.createSession).
		Doc("create a new session").
		Operation("createSession").
		Reads(TokenRequest{}).
		Writes(""))

	c.Add(ws)

	return h
}

var publicKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) {
	return PublicKey, nil
}

func (h *AuthResource) AuthorizationFilter(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {
	// auth/token is exempt
	if req.SelectedRoutePath() == "/auth/token" {
		chain.ProcessFilter(req, res)
		return
	}

	bearer := req.Request.Header.Get("Authorization")
	if strings.HasPrefix(bearer, "Bearer ") {
		token, err := jwt.Parse(bearer[7:], publicKeyFunc)
		if err == nil {
			jti := token.Claims["jti"].(string)
			if t, ok := h.store.GetSession(jti); ok {
				req.SetAttribute("token", t)
				chain.ProcessFilter(req, res)
				return
			}
		}
	}

	res.AddHeader("WWW-Authenticate", `OAuth realm="http://localhost:3001/"`)
	res.WriteErrorString(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

func (h *AuthResource) createSession(req *restful.Request, res *restful.Response) {
	tr := new(TokenRequest)
	req.ReadEntity(tr)

	if tr.Username == "username" && tr.Password == "password" {
		jti := uuid.New()
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims["jti"] = jti
		token.Claims["sub"] = tr.Username
		token.Claims["name"] = tr.Username
		token.Claims["roles"] = []string{"admin", "shmurda"}
		tokenString, err := token.SignedString(h.signingKey)
		if err != nil {
			res.WriteErrorString(500, err.Error())
		}

		h.store.SaveSession(jti, token)

		response := TokenResponse{
			AccessToken: jti,
			TokenType:   "bearer",
			ExpiresIn:   3600,
			IdToken:     tokenString,
		}

		res.WriteHeader(http.StatusFound)
		res.WriteEntity(response)
	} else {
		res.WriteErrorString(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}
}
