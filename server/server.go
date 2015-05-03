package server

import (
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful/swagger"
	"github.com/nkcraddock/gooby"
)

func BuildContainer(store gooby.Store, privateKey, publicKey []byte) *restful.Container {
	c := restful.NewContainer()

	auth := RegisterAuth(c, store, privateKey, publicKey)
	RegisterCompanies(c, store, auth)
	RegisterSwagger(c)
	RegisterStaticContent(c, "/client/build")

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

func RegisterStaticContent(container *restful.Container, root string) {
	current, _ := os.Getwd()
	staticRoot := path.Join(current, root)
	notFound := path.Join(staticRoot, "404.html")
	log.Println("Client root:", staticRoot)

	var staticHandler = func(req *restful.Request, res *restful.Response) {
		fullPath := path.Join(staticRoot, req.PathParameter("path"))
		_, err := os.Stat(fullPath)
		found := err == nil
		if found {
			http.ServeFile(res.ResponseWriter, req.Request, fullPath)
		} else {
			http.ServeFile(res.ResponseWriter, req.Request, notFound)
		}
	}

	ws := new(restful.WebService)
	ws.Route(ws.GET("/{path:*}").To(staticHandler))
	container.Add(ws)
}
