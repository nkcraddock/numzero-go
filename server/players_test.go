package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("players integration tests", func() {
	var s *ServerHarness

	BeforeEach(func() {
		authStore := numzero.NewMemoryStore()
		store := game.NewMemoryStore()
		s = NewServerHarness(authStore, store)
		s.Authenticate("username", "password")
	})

	req_chad := map[string]interface{}{
		"Name": "Chad",
	}

	Context("PUT /players", func() {
		It("adds a new player", func() {
			res := s.PUT("/players", &req_chad)
			Ω(res.Code).Should(Equal(http.StatusCreated))
		})

		It("updates an existing payer", func() {
			s.PUT("/players", &req_chad)

			req_update := map[string]interface{}{
				"Name":  "Chad",
				"Score": 1000,
			}

			s.PUT("/players", &req_update)

			p := &game.Player{}
			s.GET("/players/Chad", p)
			Ω(p.Score).Should(Equal(1000))
		})
	})

	Context("GET /players", func() {
		It("gets a player", func() {
			s.PUT("/players", &req_chad)

			p := &game.Player{}
			res := s.GET("/players/Chad", p)

			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(p.Name).Should(Equal(req_chad["Name"]))
		})
	})
})
