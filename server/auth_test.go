package server_test

import (
	"net/http"

	"github.com/nkcraddock/gooby/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AuthResource integration tests", func() {
	cfg := &server.ServerConfig{}
	s := NewServerHarness(cfg)

	Context("POST /auth/token", func() {
		It("can authenticate a username/password", func() {
			req := &server.TokenRequest{
				GrantType: "password",
				Username:  "username",
				Password:  "password",
				ClientId:  "client_id",
			}
			res := s.POST("/auth/token", &req)
			Ω(res.Code).Should(Equal(http.StatusFound))
		})
		It("wont authenticate a bad username/password", func() {
			req := &server.TokenRequest{
				GrantType: "password",
				Username:  "notarealusername",
				Password:  "nottheirrealpassword",
				ClientId:  "client_id",
			}
			res := s.POST("/auth/token", &req)
			Ω(res.Code).Should(Equal(http.StatusBadRequest))
		})
	})
})
