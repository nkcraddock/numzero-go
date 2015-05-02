package server_test

import (
	"net/http"

	"github.com/nkcraddock/gooby"
	"github.com/nkcraddock/gooby/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompanyResource integration tests", func() {
	cfg := &server.ServerConfig{}
	server := NewServerHarness(cfg)
	server.Authenticate("username", "password")

	Context("GET /companies", func() {
		It("retrieves a list of companies", func() {
			companies := []gooby.Company{}
			res := server.GET("/companies", &companies)
			Ω(companies).ShouldNot(BeEmpty())
			Ω(companies).Should(HaveLen(1))
			Ω(companies[0].Name).Should(Equal("Bloodhound Gang"))
			Ω(res.Code).Should(Equal(http.StatusOK))
		})
	})

	Context("GET /companies/{id}", func() {
		It("retrieves a single company", func() {
			company := gooby.Company{}
			res := server.GET("/companies/Bloodhound Gang", &company)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(company.Name).Should(Equal("Bloodhound Gang"))
		})
	})

	Context("POST /companies", func() {
		It("adds a new company", func() {
			company := &gooby.Company{Name: "Crips"}
			res := server.POST("/companies", &company)
			Ω(res.Code).Should(Equal(http.StatusCreated))

			newCompany := new(gooby.Company)
			res = server.GET("/companies/Crips", &newCompany)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(newCompany.Name).Should(Equal("Crips"))

			server.DELETE("/companies/Crips")
		})
	})

	Context("DELETE /companies", func() {
		It("deletes a company", func() {
			company := gooby.Company{Name: "Crips"}
			res := server.POST("/companies", &company)

			res = server.DELETE("/companies/Crips")
			Ω(res.Code).Should(Equal(http.StatusNoContent))

			res = server.GET("/companies/Crips", &company)
			Ω(res.Code).Should(Equal(http.StatusNotFound))

		})
	})
})
