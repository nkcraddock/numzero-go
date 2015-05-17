package server

import (
	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
)

func BuildContainer(store numzero.Store, gstore game.Store, privateKey, publicKey []byte, contentroot string) *restful.Container {
	c := restful.NewContainer()

	auth := RegisterAuth(c, store, privateKey, publicKey)
	RegisterRulesResource(c, gstore, auth)
	RegisterPlayersResource(c, gstore, auth)
	RegisterStaticContent(c, contentroot)

	return c
}
