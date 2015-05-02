package server_test

import (
	"bytes"
	"encoding/json"
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
}

func NewServerHarness(cfg *server.ServerConfig) *ServerHarness {
	c := server.BuildContainer(cfg)
	return &ServerHarness{container: c}
}

func request(verb, uri string, data io.Reader) *http.Request {
	req, _ := http.NewRequest(verb, uri, data)
	req.Header.Set("Authorization", "Bearer 1234")
	return req
}

func (s *ServerHarness) GET(uri string, data interface{}) (res *httptest.ResponseRecorder) {
	req := request("GET", uri, nil)
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

	req := request("POST", uri, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}

func (s *ServerHarness) DELETE(uri string) (res *httptest.ResponseRecorder) {
	req := request("DELETE", uri, nil)
	res = httptest.NewRecorder()

	s.container.ServeHTTP(res, req)

	return
}
