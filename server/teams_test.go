package server_test

import (
	"net/http"

	"github.com/nkcraddock/gooby"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TeamResource integration tests", func() {
	store := gooby.NewMemoryStore("Bloodhound Gang", "Gang of Four")
	server := NewServerHarness(store)
	server.Authenticate("username", "password")

	Context("GET /teams", func() {
		It("retrieves a list of teams", func() {
			teams := []gooby.Team{}
			res := server.GET("/teams", &teams)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(teams).ShouldNot(BeEmpty())
			Ω(teams).Should(HaveLen(2))
			Ω(teams[0].Name).Should(Equal("Bloodhound Gang"))
		})
	})

	Context("GET /teams/{id}", func() {
		It("retrieves a single team", func() {
			team := gooby.Team{}
			res := server.GET("/teams/Bloodhound Gang", &team)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(team.Name).Should(Equal("Bloodhound Gang"))
		})
	})

	Context("POST /teams", func() {
		It("adds a new team", func() {
			team := &gooby.Team{Name: "Crips"}
			res := server.POST("/teams", &team)
			Ω(res.Code).Should(Equal(http.StatusCreated))

			newTeam := new(gooby.Team)
			res = server.GET("/teams/Crips", &newTeam)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(newTeam.Name).Should(Equal("Crips"))

			server.DELETE("/teams/Crips")
		})
	})

	Context("DELETE /teams", func() {
		It("deletes a team", func() {
			team := gooby.Team{Name: "Crips"}
			res := server.POST("/teams", &team)

			res = server.DELETE("/teams/Crips")
			Ω(res.Code).Should(Equal(http.StatusNoContent))

			res = server.GET("/teams/Crips", &team)
			Ω(res.Code).Should(Equal(http.StatusNotFound))

		})
	})
})
