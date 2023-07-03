package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed frontend/dist
var frontendFolder embed.FS

func main() {
	// TODO email configuration
	// TODO discord webhook configuration
	createSampleConfig()
	config := readConfigOrExit()
	serviceMonitor := getServiceMonitorFromConfig(config)

	go monitor(&serviceMonitor)

	controller := ApiController{
		serviceMonitor: &serviceMonitor,
	}

	router := gin.Default()
	router.NoRoute(serveFiles())
	router.GET("/api", controller.indexJson)

	router.Run() // TODO add port option
}

func serveFiles() gin.HandlerFunc {
	folder, err := fs.Sub(frontendFolder, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		c.FileFromFS(path, http.FS(folder))
	}
}
