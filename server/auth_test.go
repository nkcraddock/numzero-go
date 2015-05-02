package server_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dgrijalva/jwt-go"
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
		It("returns a valid JWT", func() {
			req := &server.TokenRequest{
				GrantType: "password",
				Username:  "username",
				Password:  "password",
				ClientId:  "client_id",
			}
			res := s.POST("/auth/token", &req)
			Ω(res.Code).Should(Equal(http.StatusFound))

			tokenResponse := new(server.TokenResponse)
			body, err := ioutil.ReadAll(res.Body)
			Ω(err).ShouldNot(HaveOccurred())

			json.Unmarshal(body, &tokenResponse)
			token, err := jwt.Parse(tokenResponse.IdToken, publicKeyFunc)
			Ω(err).ShouldNot(HaveOccurred())

			Ω(token).ShouldNot(BeNil())
			Ω(token.Claims["sub"]).Should(Equal("username"))
		})
	})
})

var publicKeyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) {
	return server.PublicKey, nil
}
