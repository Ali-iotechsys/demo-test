package demo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Adder", func() {

	// Create a Resty Client
	client := resty.New()
	resp, err := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

	var startStr = "start"
	var endStr = "end"

	BeforeEach(func() {
		fmt.Println(startStr)
	})

	AfterEach(func() {
		fmt.Println(endStr)
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
