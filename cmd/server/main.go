package main

import (
	"log"
	"net/http"

	"github.com/nkcraddock/gooby/server"
)

func main() {
	cfg := &server.ServerConfig{Addr: ":3001"}
	c := server.BuildContainer(cfg)

	server := &http.Server{Addr: cfg.Addr, Handler: c}
	log.Fatal(server.ListenAndServe())
}
