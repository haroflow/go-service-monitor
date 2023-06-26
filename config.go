package main

type Config struct {
	Monitors MonitorsConfig `json:"monitors"`
}

type MonitorsConfig struct {
	HTTP []HTTPConfig `json:"http"`
	TCP  []TCPConfig  `json:"tcp"`
	DNS  []DNSConfig  `json:"dns"`
}

type HTTPConfig struct {
	Address string `json:"address"`
	Timeout int    `json:"timeout"`
	// Interval int
}

type TCPConfig struct {
	Address string `json:"address"`
	Port    int16  `json:"port"`
	Timeout int    `json:"timeout"`
	// Interval int
}

type DNSConfig struct {
	Address string `json:"address"`
	// Interval int
}

func newSampleConfig() Config {
	return Config{
		Monitors: MonitorsConfig{
			HTTP: []HTTPConfig{
				{Address: "http://www.google.com", Timeout: 15},
				{Address: "https://www.google.com", Timeout: 15},
				{Address: "https://www.addressthatdoesnotexist.com", Timeout: 15},
			},
			TCP: []TCPConfig{
				{Address: "www.google.com", Port: 80, Timeout: 15},
				{Address: "www.google.com", Port: 443, Timeout: 15},
				{Address: "www.google.com", Port: 1433, Timeout: 15},
				{Address: "1.1.1.1", Port: 53, Timeout: 15},
			},
			DNS: []DNSConfig{
				{Address: "google.com"},
				{Address: "amazon.com"},
				{Address: "whois.com"},
			},
		},
	}
}
