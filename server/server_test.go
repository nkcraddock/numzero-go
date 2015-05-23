package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero"
	"github.com/nkcraddock/numzero/game"
	"github.com/nkcraddock/numzero/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("server integration tests", func() {
	var s *ServerHarness

	req_user := map[string]interface{}{"name": "shmurda"}

	req_coffee := map[string]interface{}{
		"code":   "coffee",
		"desc":   "made coffee",
		"points": 5,
	}

	BeforeEach(func() {
		authStore := numzero.NewMemoryStore()
		store, err := game.NewRedisStore("localhost:6379", "", 10)
		Ω(err).ShouldNot(HaveOccurred())
		store.FlushDb()
		s = NewServerHarness(authStore, store)
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
			Ω(rule.Points).Should(Equal(5))
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

	Context("/players/{player}/activites", func() {
		req_tooMuchCoffee := server.Activity{
			Description: "SHmurda",
			Url:         "",
			Scores: map[string]int{
				"coffee": 1000,
			},
		}

		It("stores an event for the player", func() {
			s.PUT("/players", &req_user)
			s.PUT("/rules", &req_coffee)

			res := s.POST("/players/shmurda/activities", req_tooMuchCoffee)
			Ω(res.Code).Should(Equal(http.StatusOK))
		})
	})
})
