package main

import (
	"log"
	"time"

	"github.com/haroflow/go-service-monitor/lib"
)

type ServiceMonitor struct {
	// TODO Mutex? RWMutex? Save directly to database?
	LastCheck  time.Time
	HTTPChecks []HTTPCheck
	TCPChecks  []TCPCheck
	DNSChecks  []DNSCheck
}

type HTTPCheck struct {
	Address string
	Status  bool
	Timeout int
	// ExpectedStatusCode int
}

type TCPCheck struct {
	Address string
	Port    int16
	Status  bool
	Timeout int
}

type DNSCheck struct {
	Address string
	Status  bool
	// Server string
}

func getServiceMonitorFromConfig(config Config) ServiceMonitor {
	serviceMonitor := ServiceMonitor{}

	for _, http := range config.Monitors.HTTP {
		var httpCheck = HTTPCheck{
			Address: http.Address,
			Timeout: http.Timeout,
			Status:  false,
		}
		serviceMonitor.HTTPChecks = append(serviceMonitor.HTTPChecks, httpCheck)
	}
	for _, tcp := range config.Monitors.TCP {
		var tcpCheck = TCPCheck{
			Address: tcp.Address,
			Port:    tcp.Port,
			Timeout: tcp.Timeout,
			Status:  false,
		}
		serviceMonitor.TCPChecks = append(serviceMonitor.TCPChecks, tcpCheck)
	}
	for _, dns := range config.Monitors.DNS {
		var dnsCheck = DNSCheck{
			Address: dns.Address,
			Status:  false,
		}
		serviceMonitor.DNSChecks = append(serviceMonitor.DNSChecks, dnsCheck)
	}

	return serviceMonitor
}

func monitor(serviceMonitor *ServiceMonitor) { // TODO parallelize
	for {
		serviceMonitor.LastCheck = time.Now()

		for i := range serviceMonitor.HTTPChecks {
			http := &serviceMonitor.HTTPChecks[i]
			err := lib.TestHTTPEndpoint(http.Address, http.Timeout)
			if err != nil {
				http.Status = false
				log.Printf("HTTP FAIL | %+v | ERR: %v\n", http, err)
			} else {
				http.Status = true
				log.Printf("HTTP OK | %+v", http)
			}
		}

		for i := range serviceMonitor.TCPChecks {
			tcp := &serviceMonitor.TCPChecks[i]
			err := lib.TestTCPEndpoint(tcp.Address, tcp.Port, tcp.Timeout)
			if err != nil {
				tcp.Status = false
				log.Printf("TCP FAIL | %+v | ERR: %v\n", tcp, err)
			} else {
				tcp.Status = true
				log.Printf("TCP OK | %+v", tcp)
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
