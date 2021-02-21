package demo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("Demo Test", func() {

	// Create a Resty Client
	client := resty.New()

	// Test connection
	_, err := client.R().Get("")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	BeforeEach(func() {
		fmt.Println("this is start test")
	})

	AfterEach(func() {
		fmt.Println("this is end test")
	})

	Context("when sum are positive", func() {

		It("adds two numbers", func() {
			sum, err := Add(2, 3)
			Expect(err).NotTo(HaveOccurred())
			Expect(sum).To(Equal(5))
		})

	})

	Context("when sum is negative", func() {

		It("returns an err", func() {
			_, err := Add(-1, -1)
			Expect(err).To(HaveOccurred())
		})
	})

})
