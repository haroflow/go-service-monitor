package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO email configuration
	// TODO discord webhook configuration

	serviceMonitor = ServiceMonitor{ // TODO Should be loaded from file
		HTTPChecks: []HTTPCheck{
			{Address: "http://www.google.com", Status: false},
			{Address: "https://www.google.com", Status: false},
			{Address: "https://www.addressthatdoesnotexist.com", Status: false},
		},
		TCPChecks: []TCPCheck{
			{Address: "www.google.com", Port: 80, Status: false},
			{Address: "www.google.com", Port: 443, Status: false},
			{Address: "www.google.com", Port: 1433, Status: false},
			{Address: "1.1.1.1", Port: 53, Status: false},
		},
		DNSChecks: []DNSCheck{
			{Address: "google.com", Status: false},
			{Address: "amazon.com", Status: false},
			{Address: "whois.com", Status: false},
		},
	}

	go monitor()

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", index)
	router.GET("/json", indexJson)
	router.Run() // TODO add port option
}
