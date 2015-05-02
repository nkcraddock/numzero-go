package server_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/emicklei/go-restful"
	"github.com/nkcraddock/gooby/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Server Suite")
}

type ServerHarness struct {
	container *restful.Container
	token     *string
}

func NewServerHarness(cfg *server.ServerConfig) *ServerHarness {
	c := server.BuildContainer(cfg)
	return &ServerHarness{container: c}
}

func (s *ServerHarness) request(verb, uri string, data io.Reader) *http.Request {
	req, _ := http.NewRequest(verb, uri, data)

	if s.token != nil {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *s.token))
	}

	return req
}

func (s *ServerHarness) Authenticate(username, password string) error {
	req := &server.TokenRequest{
		GrantType: "password",
		Username:  "username",
		Password:  "password",
		ClientId:  "client_id",
	}
	res := s.POST("/auth/token", &req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	tr := new(server.TokenResponse)

	json.Unmarshal(body, tr)

	s.token = &tr.IdToken

	return nil
}

func (s *ServerHarness) GET(uri string, data interface{}) (res *httptest.ResponseRecorder) {
	req := s.request("GET", uri, nil)
	req.Header.Set("Accept", "application/json")
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal(body, &data)

	return
}

func (s *ServerHarness) POST(uri string, postdata interface{}) (res *httptest.ResponseRecorder) {
	data, err := json.Marshal(postdata)
	if err != nil {
		return
	}

	req := s.request("POST", uri, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}

func (s *ServerHarness) DELETE(uri string) (res *httptest.ResponseRecorder) {
	req := s.request("DELETE", uri, nil)
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}
