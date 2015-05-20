package server

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/emicklei/go-restful"
	"github.com/gorilla/schema"
	"github.com/nkcraddock/numzero/game"
)

type IntegrationsResource struct {
	store game.Store
}

func RegisterIntegrationResource(c *restful.Container, store game.Store, auth *AuthResource) *IntegrationsResource {
	h := &IntegrationsResource{store: store}

	ws := new(restful.WebService)

	ws.Path("/integrations").
		Doc("Handle data from webhooks and such").
		Consumes(restful.MIME_XML, restful.MIME_JSON, "application/x-www-form-urlencoded").
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.POST("/slack").To(h.slack).
		Doc("Post data from a slack outgoing webhook").
		Operation("slack").
		Reads(slackdata{}))

	c.Add(ws)

	return h
}

func (h *IntegrationsResource) slack(req *restful.Request, res *restful.Response) {
	err := req.Request.ParseForm()
	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, "")
		log.Println("Bad slack msg", err)
		return
	}

	d := new(slackdata)

	decoder := schema.NewDecoder()
	err = decoder.Decode(d, req.Request.PostForm)

	if err != nil {
		res.WriteErrorString(http.StatusBadRequest, "")
		log.Println("Bad slack msg", err)
		return
	}

	if err = h.processSlack(d); err != nil {
		res.WriteErrorString(http.StatusBadRequest, "")
		log.Println("Bad slack data", err)
		return
	}

	response := map[string]string{
		"text": "murrh",
	}

	res.WriteEntity(response)
}

func (h *IntegrationsResource) processSlack(data *slackdata) error {
	if data.Token != config.SlackToken {
		return errors.New("Invalid slack token")
	}

	p, err := h.store.GetPlayer(data.User_name)
	if err != nil {
		p = game.NewPlayer(data.User_name)
		if err := h.store.SavePlayer(p); err != nil {
			return err
		}
	}

	if scores := h.scoreSlack(data); scores != nil {
		event := &game.Event{
			Description: "Said something on slack",
			Url:         "https://www.youtube.com/watch?v=hpigjnKl7nI",
			Scores:      scores,
		}

		if err := p.AddEvent(event); err != nil {
			return err
		}

		h.store.SavePlayer(p)
	}

	return nil
}

func (h *IntegrationsResource) scoreSlack(data *slackdata) []game.Score {
	utterances := 0
	for _, w := range strings.Split(data.Text, " ") {
		word := strings.ToLower(w)
		if _, ok := words[word]; ok {
			words[word] += 1
			utterances += 1
		}
	}

	if utterances > 0 {
		profanity, err := h.store.GetRule("slack:profanity")
		if err != nil {
			return nil
		}

		return []game.Score{
			game.Score{&profanity, utterances},
		}
	}

	return nil
}

type slackdata struct {
	Token        string
	Team_id      string
	Team_domain  string
	Channel_id   string
	Channel_name string
	Timestamp    float32
	User_id      string
	User_name    string
	Text         string
	Trigger_word string
	Service_id   int
}

var words map[string]int = map[string]int{"i feel like": 0}
