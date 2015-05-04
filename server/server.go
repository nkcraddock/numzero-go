package server

import (
	"os"
	"path/filepath"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/nkcraddock/gooby"
)

func BuildContainer(store gooby.Store, privateKey, publicKey []byte, contentroot string) *restful.Container {
	c := restful.NewContainer()

	auth := RegisterAuth(c, store, privateKey, publicKey)
	RegisterCompanies(c, store, auth)
	RegisterSwagger(c)
	RegisterStaticContent(c, contentroot)

	return c
}

func RegisterSwagger(container *restful.Container) {
	current, _ := os.Getwd()
	config := swagger.Config{
		WebServices:     container.RegisteredWebServices(),
		ApiPath:         "/apidocs.json",
		SwaggerPath:     "/apidocs/",
		SwaggerFilePath: filepath.Join(current, "swagger"),
	}

	swagger.RegisterSwaggerService(config, container)
}
