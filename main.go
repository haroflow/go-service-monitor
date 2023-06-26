package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const configPath string = "config.json"
const sampleConfigPath string = "sample.config.json"

func main() {
	// TODO email configuration
	// TODO discord webhook configuration
	createSampleConfig()
	config := readConfigOrExit()
	serviceMonitor := getServiceMonitorFromConfig(config)

	go monitor(&serviceMonitor)

	controller := IndexController{
		serviceMonitor: serviceMonitor,
	}

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controller.index)
	router.GET("/json", controller.indexJson)
	router.Run() // TODO add port option
}

func createSampleConfig() {
	file, err := os.Create(sampleConfigPath)
	if err != nil {
		log.Printf("Cannot create %v: %v\n", sampleConfigPath, err)
		return
	}
	defer file.Close()

	sampleConfig := newSampleConfig()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(sampleConfig)
	if err == nil {
		log.Printf("Cannot write to %v: %v\n", sampleConfigPath, err)
		return
	}
}

func readConfigOrExit() Config {
	_, err := os.Stat(configPath)
	if err != nil {
		log.Printf(
			"Please configure monitoring in %v. You can use %v as reference. Error: %v\n",
			configPath, sampleConfigPath, err)
		os.Exit(1)
	}

	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("Cannot read %v: %v\n", configPath, err)
		os.Exit(1)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Printf("Cannot read %v as JSON: %v\n", configPath, err)
		os.Exit(1)
	}

	return config
}
