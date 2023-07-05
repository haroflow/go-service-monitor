package lib

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/miekg/dns"
)

// CheckHTTPEndpoint sends a GET request to the specified address, with a
// default 10 second timeout, and expects a 200 OK response.
func CheckHTTPEndpoint(address string, timeoutSeconds int) error {
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return fmt.Errorf("address should begin with http:// or https://")
	}

	timeout := time.Duration(timeoutSeconds) * time.Second
	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(address)
	if err != nil {
		return fmt.Errorf("error on GET request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return fmt.Errorf("request to %v returned status %v", address, res.StatusCode)
	}

	return nil
}

// CheckTCPEndpoint sends a TCP request to the specified address and port,
// with a default timeout of 10 seconds.
func CheckTCPEndpoint(address string, port int16, timeoutSeconds int) error {
	timeout := time.Duration(timeoutSeconds) * time.Second
	dialer := net.Dialer{
		Timeout: timeout,
	}

	conn, err := dialer.Dial("tcp", fmt.Sprintf("%v:%v", address, port))
	if err != nil {
		return fmt.Errorf("error on TCP connection: %w", err)
	}
	defer conn.Close()

	return nil
}

// CheckDNSResponse requests an IP lookup for the address.
// You may also specify a DNS server address, or pass an empty string to
// use the local resolver.
func CheckDNSResponse(address string, server string) error {
	if server == "" {
		_, err := net.LookupIP(address)
		return err
	} else {
		c := dns.Client{}
		m := dns.Msg{}
		m.SetQuestion(address+".", dns.TypeA)
		r, _, err := c.Exchange(&m, server+":53")
		if err != nil {
			return err
		}
		if len(r.Answer) == 0 {
			return errors.New("dns query: no results")
		}
		return nil
	}
}
