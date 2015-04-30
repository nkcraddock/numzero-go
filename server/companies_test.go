package server_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CompanyResource integration tests", func() {
	Context("GET /companies", func() {
		It("retrieves a list of companies", func() {
			Ω("").Should(Equal(""))
		})
	})
})
