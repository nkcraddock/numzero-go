package main

import (
	"log"
	"net/http"

	"github.com/nkcraddock/gooby"
	"github.com/nkcraddock/gooby/server"
)

func main() {
	addr := ":3001"
	store := gooby.NewStore()
	c := server.BuildContainer(store)

	server := &http.Server{Addr: addr, Handler: c}
	log.Fatal(server.ListenAndServe())
}
