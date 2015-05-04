package server

import (
	"log"
	"mime"
	"os"
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
	RegisterStaticContent(c)

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

func RegisterStaticContent(container *restful.Container) {
	notFound, _ := gooby.Asset("404.html")

	var staticHandler = func(req *restful.Request, res *restful.Response) {
		filePath := req.PathParameter("path")

		if filePath == "" {
			filePath = "index.html"
		}

		if data, err := gooby.Asset(filePath); err == nil {
			mimetype := mime.TypeByExtension(filepath.Ext(filePath))
			res.AddHeader("Content-Type", mimetype)
			res.Write(data)
		} else {
			log.Println("Not found:", filePath, gooby.AssetNames())
			res.Write(notFound)
		}
	}

	ws := new(restful.WebService)
	ws.Route(ws.GET("/{path:*}").To(staticHandler))
	container.Add(ws)
}
