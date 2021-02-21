package demo

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Demo Test", func() {

	// Create an API Client
	client := resty.New()

	// Init authentication token
	authResult := struct {
		Jwt string `json:"jwt"`
	}{Jwt: ""}

	BeforeEach(func() {
		fmt.Println("this is start test")
	})

	AfterEach(func() {
		fmt.Println("this is end test")
	})

	Context("Test Connectivity Feature", func() {

		It("EB Controller API is reachable", func() {
			resp, err := client.R().Get("http://192.168.0.7:8085/api/")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})

		It("EB Controller can authenticate a user", func() {
			resp, err := client.R().
				SetHeader("accept", "application/json").
				SetHeader("Content-Type", "application/json").
				SetBody(`{"Username": "iotech",  "Password": "EdgeBuilder123"}`).
				SetResult(&authResult).
				Post("http://192.168.0.7:8085/api/auth")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(authResult.Jwt).NotTo(BeEmpty())
		})

	})

	Context("Test Groups Feature", func() {

		It("EB Controller API can add a group", func() {
			resp, err := client.R().
				SetHeader("accept", "application/json").
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", authResult.Jwt).
				SetBody(`[{"Name":"Group1","Description":"This is Group 1"}]`).
				Post("http://192.168.0.7:8085/api/groups")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})

		It("EB Controller API can view all groups", func() {
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", authResult.Jwt).
				Get("http://192.168.0.7:8085/api/groups")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})

		It("EB Controller API can view a group by Name", func() {
			resp, err := client.R().
				SetQueryParams(map[string]string{
					"group": "Group1",
				}).
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", authResult.Jwt).
				Get("http://192.168.0.7:8085/api/groups")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})

		It("EB Controller API can remove a group", func() {
			deleteResult := struct {
				Deleted int `json:"Deleted"`
			}{Deleted: 0}
			resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetHeader("Authorization", authResult.Jwt).
				SetBody(`["Group1"]`).
				SetResult(&deleteResult).
				Delete("http://192.168.0.7:8085/api/groups")
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			Expect(deleteResult.Deleted).To(Equal(1))
		})

	})

})
