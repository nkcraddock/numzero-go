package server_test

import (
	"net/http"

	"github.com/nkcraddock/gooby/server"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompanyResource integration tests", func() {
	cfg := &server.ServerConfig{}
	server := NewServerHarness(cfg)

	Context("GET /companies", func() {
		It("retrieves a list of companies", func() {
			companies := []string{}
			res := server.GET("/companies", &companies)
			Ω(companies).ShouldNot(BeEmpty())
			Ω(companies).Should(HaveLen(2))
			Ω(res.Code).Should(Equal(http.StatusOK))
		})
	})
})
