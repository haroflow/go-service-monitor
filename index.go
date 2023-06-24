package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type indexVM struct {
	LastCheck time.Time
	Services  []serviceStatusVM
}

type serviceStatusVM struct {
	CheckType string
	Address   string
	Status    bool
}

func newIndexViewModel(sm ServiceMonitor) []serviceStatusVM {
	var vm []serviceStatusVM

	for _, h := range sm.HTTPChecks {
		vm = append(vm, serviceStatusVM{
			CheckType: "HTTP/HTTPS",
			Address:   h.Address,
			Status:    h.Status,
		})
	}
	for _, t := range sm.TCPChecks {
		vm = append(vm, serviceStatusVM{
			CheckType: fmt.Sprintf("TCP %v", t.Port),
			Address:   t.Address,
			Status:    t.Status,
		})
	}
	for _, h := range sm.DNSChecks {
		vm = append(vm, serviceStatusVM{
			CheckType: "DNS",
			Address:   h.Address,
			Status:    h.Status,
		})
	}

	return vm
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", indexVM{
		LastCheck: serviceMonitor.LastCheck,
		Services:  newIndexViewModel(serviceMonitor),
	})
}

func indexJson(c *gin.Context) {
	c.IndentedJSON(200, serviceMonitor)
}
