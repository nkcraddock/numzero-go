package server_test

import (
	"net/http"

	"github.com/nkcraddock/numzero"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("rawResource integration tests", func() {
	store := numzero.NewMemoryStore("Bloodhound Gang", "Gang of Four")
	server := NewServerHarness(store)
	server.Authenticate("username", "password")

	Context("GET /raw", func() {
		It("retrieves a list of raw", func() {
			raw := []numzero.RawBuildData{}
			res := server.GET("/raw", &raw)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(raw).ShouldNot(BeEmpty())
			Ω(raw).Should(HaveLen(2))
			Ω(raw[0].Name).Should(Equal("Bloodhound Gang"))
		})
	})

	Context("GET /raw/{id}", func() {
		It("retrieves a single raw", func() {
			raw := numzero.RawBuildData{}
			res := server.GET("/raw/Bloodhound Gang", &raw)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(raw.Name).Should(Equal("Bloodhound Gang"))
		})
	})

	Context("POST /raw", func() {
		It("adds a new raw", func() {
			raw := &numzero.RawBuildData{Name: "Crips"}
			res := server.POST("/raw", &raw)
			Ω(res.Code).Should(Equal(http.StatusCreated))

			newraw := new(numzero.RawBuildData)
			res = server.GET("/raw/Crips", &newraw)
			Ω(res.Code).Should(Equal(http.StatusOK))
			Ω(newraw.Name).Should(Equal("Crips"))

			server.DELETE("/raw/Crips")
		})
	})

	Context("DELETE /raw", func() {
		It("deletes a raw", func() {
			raw := numzero.RawBuildData{Name: "Crips"}
			res := server.POST("/raw", &raw)

			res = server.DELETE("/raw/Crips")
			Ω(res.Code).Should(Equal(http.StatusNoContent))

			res = server.GET("/raw/Crips", &raw)
			Ω(res.Code).Should(Equal(http.StatusNotFound))

		})
	})
})
