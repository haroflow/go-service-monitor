package main

import (
	"encoding/json"
	"log"
	"os"
)

const configPath string = "config.json"
const sampleConfigPath string = "sample.config.json"

type Config struct {
	Monitors MonitorsConfig `json:"monitors"`
}

type MonitorsConfig struct {
	HTTP []HTTPConfig `json:"http"`
	TCP  []TCPConfig  `json:"tcp"`
	DNS  []DNSConfig  `json:"dns"`
}

type HTTPConfig struct {
	DisplayName string `json:"displayName"`
	Address     string `json:"address"`
	Timeout     int    `json:"timeout"`
	// Interval int
}

type TCPConfig struct {
	DisplayName string `json:"displayName"`
	Address     string `json:"address"`
	Port        int16  `json:"port"`
	Timeout     int    `json:"timeout"`
	// Interval int
}

type DNSConfig struct {
	DisplayName string `json:"displayName"`
	Address     string `json:"address"`
	// Interval int
}

func newSampleConfig() Config {
	return Config{
		Monitors: MonitorsConfig{
			HTTP: []HTTPConfig{
				{DisplayName: "Google HTTP", Address: "http://www.google.com", Timeout: 15},
				{DisplayName: "Google HTTPS", Address: "https://www.google.com", Timeout: 15},
				{DisplayName: "Invalid Site HTTPS", Address: "https://www.addressthatdoesnotexist.com", Timeout: 15},
			},
			TCP: []TCPConfig{
				{DisplayName: "Google Port 80", Address: "www.google.com", Port: 80, Timeout: 15},
				{DisplayName: "Google Port 443", Address: "www.google.com", Port: 443, Timeout: 15},
				{DisplayName: "Google SQL Server Port 1433", Address: "www.google.com", Port: 1433, Timeout: 15},
				{DisplayName: "Cloudflare DNS Port 53", Address: "1.1.1.1", Port: 53, Timeout: 15},
			},
			DNS: []DNSConfig{
				{DisplayName: "Google A Record", Address: "google.com"},
				{DisplayName: "Amazon A Record", Address: "amazon.com"},
				{DisplayName: "Whois A Record", Address: "whois.com"},
			},
		},
	}
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
	if err != nil {
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
