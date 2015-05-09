package server

import (
	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
)

func BuildContainer(store numzero.Store, privateKey, publicKey []byte, contentroot string) *restful.Container {
	c := restful.NewContainer()

	auth := RegisterAuth(c, store, privateKey, publicKey)
	RegisterTeams(c, store, auth)
	RegisterStaticContent(c, contentroot)

	return c
}
