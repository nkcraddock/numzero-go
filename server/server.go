package server

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
)

var config appconfig

func BuildContainer(store numzero.Store, gstore game.Store, privateKey, publicKey []byte, contentroot string) *restful.Container {
	config = appconfig{
		SlackToken: "50u6HWjJmjiK0dL9ViWKXPSu",
	}

	c := restful.NewContainer()
	c.EnableContentEncoding(true)

	auth := RegisterAuth(c, store, privateKey, publicKey)
	RegisterRulesResource(c, gstore, auth)
	RegisterPlayersResource(c, gstore, auth)
	RegisterEventsResource(c, gstore, auth)
	RegisterStaticContent(c, contentroot)

	return c
}

type appconfig struct {
	SlackToken string
}

func handleError(err error, msg string, status int, res *restful.Response) bool {
	if err == nil {
		return false
	}

	if msg == "" {
		msg = http.StatusText(status)
	}

	res.WriteHeader(status)
	response := map[string]interface{}{
		"error": msg,
	}
	log.Println("ERROR", err)
	res.WriteEntity(response)
	return true
}
