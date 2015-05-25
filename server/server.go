package server

import (
	"log"
	"net/http"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
)

func BuildContainer(store numzero.Store, gstore game.Store, cfg ServerConfig) *restful.Container {
	c := restful.NewContainer()
	c.EnableContentEncoding(true)

	auth := RegisterAuth(c, store, cfg.PrivateKey, cfg.PublicKey)
	RegisterRulesResource(cfg.RootApiPath, c, gstore, auth)
	RegisterPlayersResource(cfg.RootApiPath, c, gstore, auth)
	RegisterEventsResource(cfg.RootApiPath, c, gstore, auth, cfg.WebhookUrl)
	RegisterStaticContent(c, cfg.ContentRoot)

	return c
}

type ServerConfig struct {
	PrivateKey  []byte
	PublicKey   []byte
	ContentRoot string
	WebhookUrl  string
	RootApiPath string
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
