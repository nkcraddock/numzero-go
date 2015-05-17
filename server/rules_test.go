package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rules integration tests", func() {
	var s *ServerHarness

	BeforeEach(func() {
		authStore := numzero.NewMemoryStore()
		store := game.NewMemoryStore()
		s = NewServerHarness(authStore, store)
		s.Authenticate("username", "password")
	})

	req_coffee := map[string]interface{}{
		"code":   "coffee",
		"desc":   "Made a new pot of coffee",
		"points": 1,
	}

	Context("PUT /rules", func() {
		It("adds a new rule", func() {
			res := s.PUT("/rules", &req_coffee)
			Ω(res.Code).Should(Equal(http.StatusCreated))
		})

		It("updates an existing rule", func() {
			s.PUT("/rules", &req_coffee)

			modified_coffee := map[string]interface{}{
				"code":   "coffee",
				"desc":   "You kill the joe you make some mo'",
				"points": 10,
			}

			s.PUT("/rules", &modified_coffee)

			rule := game.Rule{}
			s.GET("/rules/coffee", &rule)
			Ω(rule.Description).Should(Equal(modified_coffee["desc"]))
			Ω(rule.Points).Should(Equal(modified_coffee["points"]))
		})
	})

	Context("GET /rules", func() {
		It("retrieves a rule", func() {
			s.PUT("/rules", &req_coffee)

			rule := game.Rule{}
			res := s.GET("/rules/coffee", &rule)

			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(rule.Description).Should(Equal(req_coffee["desc"]))
		})
	})
})
