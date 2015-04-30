package server_test

import (
	"bytes"
	"encoding/json"
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
}

func NewServerHarness(cfg *server.ServerConfig) *ServerHarness {
	c := server.BuildContainer(cfg)
	return &ServerHarness{container: c}
}

func (s *ServerHarness) GET(uri string, data interface{}) (res *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", uri, nil)
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

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}

func (s *ServerHarness) DELETE(uri string) (res *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("DELETE", uri, nil)
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}
