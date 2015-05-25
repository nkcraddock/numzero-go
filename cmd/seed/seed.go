package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/nkcraddock/numzero/game"
)

func main() {
	var path string

	flag.StringVar(&path, "path", "testdata/seed.json", "-path path_to_seed.json")
	flag.Parse()

	data, _ := ioutil.ReadFile(path)
	var bags map[string][]map[string]interface{}
	if err := json.Unmarshal(data, &bags); err != nil {
		log.Println("Error unmarshalling:", err)
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	if players, ok := bags["players"]; ok {
		for _, p := range players {
			req := request("PUT", "/players", p)
			if _, err := client.Do(req); err != nil {
				log.Println("ERROR", err)
			}
		}
	}

	if rules, ok := bags["rules"]; ok {
		for _, r := range rules {
			req := request("PUT", "/rules", r)
			if _, err := client.Do(req); err != nil {
				log.Println("ERROR", err)
			}
		}
	}

	if events, ok := bags["events"]; ok {
		for _, r := range events {
			req := request("POST", "/events", r)
			if _, err := client.Do(req); err != nil {
				log.Println("ERROR", err)
			}
		}
	}
}

func request(verb, path string, thing interface{}) *http.Request {
	url := fmt.Sprintf("https://localhost:3001%s", path)
	pJson, _ := json.Marshal(thing)
	req, _ := http.NewRequest(verb, url, bytes.NewBuffer(pJson))
	req.Header.Set("content-type", "application/json")
	return req
}

type seedData struct {
	players []game.Player
	rules   []game.Rule
	events  []game.Event
}
