package main

import (
	"log"
	"time"

	"github.com/haroflow/go-service-monitor/lib"
)

var serviceMonitor ServiceMonitor

type ServiceMonitor struct {
	// TODO Mutex? RWMutex? Save directly to database?
	LastCheck  time.Time
	HTTPChecks []HTTPCheck
	TCPChecks  []TCPCheck
	DNSChecks  []DNSCheck
}

type HTTPCheck struct {
	Address string
	Status  bool // TODO split configuration from execution status
	// Timeout int
	// ExpectedStatusCode int
}

type TCPCheck struct {
	Address string
	Port    int16
	Status  bool // TODO split configuration from execution status
	// Timeout int
}

type DNSCheck struct {
	Address string
	Status  bool // TODO split configuration from execution status
	// Server string
}

func monitor() { // TODO parallelize
	for {
		serviceMonitor.LastCheck = time.Now()

		for i := range serviceMonitor.HTTPChecks {
			httpCheck := &serviceMonitor.HTTPChecks[i]
			err := lib.TestHTTPEndpoint(httpCheck.Address)
			if err != nil {
				httpCheck.Status = false
				log.Printf("HTTP FAIL | %+v | ERR: %v\n", httpCheck, err)
			} else {
				httpCheck.Status = true
				log.Printf("HTTP OK | %+v", httpCheck)
			}
		}

		for i := range serviceMonitor.TCPChecks {
			tcpCheck := &serviceMonitor.TCPChecks[i]
			err := lib.TestTCPEndpoint(tcpCheck.Address, tcpCheck.Port)
			if err != nil {
				tcpCheck.Status = false
				log.Printf("TCP FAIL | %+v | ERR: %v\n", tcpCheck, err)
			} else {
				tcpCheck.Status = true
				log.Printf("TCP OK | %+v", tcpCheck)
			}
		}

		for i := range serviceMonitor.DNSChecks {
			dnsCheck := &serviceMonitor.DNSChecks[i]
			err := lib.TestDNSResponse(dnsCheck.Address)
			if err != nil {
				dnsCheck.Status = false
				log.Printf("DNS FAIL | %+v | ERR: %v\n", dnsCheck, err)
			} else {
				dnsCheck.Status = true
				log.Printf("DNS OK | %+v", dnsCheck)
			}
		}

		time.Sleep(5 * time.Second)
	}
}
