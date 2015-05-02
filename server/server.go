package server

import (
	"os"
	"path/filepath"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
)

type ServerConfig struct {
	Addr string
}

func BuildContainer(cfg *ServerConfig) *restful.Container {
	c := restful.NewContainer()

	RegisterCompanies(c)
	RegisterSwagger(c)

	c.Filter(Authenticate)

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
