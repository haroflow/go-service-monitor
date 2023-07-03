package main

import (
	"github.com/gin-gonic/gin"
)

type ApiController struct {
	serviceMonitor *ServiceMonitor
}

func (ctrl *ApiController) indexJson(c *gin.Context) {
	c.JSON(200, ctrl.serviceMonitor) // FIXME create view model, shouldn't send internal configurations to frontend.
}
