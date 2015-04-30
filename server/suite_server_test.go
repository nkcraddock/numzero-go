package server_test

import (
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
