package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("players integration tests", func() {
	var s *ServerHarness

	BeforeEach(func() {
		s = NewServerHarness()
		s.GameStore.SaveRule(&game.Rule{"coffee", "made coffee", 1})
		s.GameStore.SaveRule(&game.Rule{"highfive", "high-fived someone", -10})

		s.Authenticate("username", "password")
	})

	req_chad := map[string]interface{}{"Name": "Chad"}
	req_roger := map[string]interface{}{"Name": "Roger"}

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

		It("gets a list of players", func() {
			s.PUT("/players", &req_chad)
			s.PUT("/players", &req_roger)
			names := []string{"Roger", "Chad"}

			var players []game.Player

			res := s.GET("/players", &players)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(players).Should(HaveLen(2))
			Ω(names).Should(ContainElement(players[0].Name))
			Ω(names).Should(ContainElement(players[1].Name))
		})
	})
})
