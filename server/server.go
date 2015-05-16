package server

import (
	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/numzero"
)

func BuildContainer(store numzero.Store, privateKey, publicKey []byte, contentroot string) *restful.Container {
	c := restful.NewContainer()

	RegisterAuth(c, store, privateKey, publicKey)
	RegisterStaticContent(c, contentroot)

	return c
}
