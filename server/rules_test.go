package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero/game"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rules integration tests", func() {
	var s *ServerHarness

	req_coffee := map[string]interface{}{
		"code":   "coffee",
		"desc":   "Made a new pot of coffee",
		"points": 1,
	}

	BeforeEach(func() {
		s = NewServerHarness()
		s.Authenticate("username", "password")
	})

	Context("/rules", func() {
		It("gets a list of rules", func() {
			res := s.PUT("/rules", &req_coffee)

			var results []game.Rule
			res = s.GET("/rules", &results)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(results).Should(HaveLen(1))
		})

		It("gets an existing rule by code", func() {
			res := s.PUT("/rules", &req_coffee)

			var rule game.Rule
			res = s.GET("/rules/coffee", &rule)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(rule.Points).Should(Equal(1))
		})

		It("adds a new rule", func() {
			res := s.PUT("/rules", &req_coffee)
			Ω(res.Code).Should(Equal(http.StatusCreated))
		})

		It("updates an existing rule", func() {
			s.PUT("/rules", &req_coffee)

			req_modified := map[string]interface{}{
				"code":   "coffee",
				"desc":   "talked about coffee",
				"points": 1,
			}

			res := s.PUT("/rules", &req_modified)
			Ω(res.Code).Should(Equal(http.StatusCreated))

			var rule game.Rule
			res = s.GET("/rules/coffee", &rule)
			Ω(res.Code).Should(Equal(http.StatusOK))

			Ω(rule.Description).Should(Equal("talked about coffee"))
			Ω(rule.Points).Should(Equal(1))
		})
	})

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

		It("can have negative points", func() {
			req_rule := map[string]interface{}{
				"code":   "highfive",
				"desc":   "high-fived someone",
				"points": -1,
			}

			res := s.PUT("/rules", &req_rule)
			Ω(res.Code).Should(Equal(http.StatusCreated))

			rule := game.Rule{}
			s.GET("/rules/highfive", &rule)
			Ω(rule.Points).Should(Equal(-1))
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
