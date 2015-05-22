package server

import (
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
	RegisterStaticContent(c, contentroot)

	return c
}

type appconfig struct {
	SlackToken string
}
