package server

import (
	"net/http"

	"github.com/emicklei/go-restful"
)

func Authenticate(req *restful.Request, res *restful.Response, chain *restful.FilterChain) {
	token := req.Request.Header.Get("Authorization")

	if len(token) == 0 || token != "Bearer 1234" {
		res.AddHeader("WWW-Authenticate", `OAuth realm="http://localhost:3001/"`)
		res.WriteErrorString(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	chain.ProcessFilter(req, res)
}
