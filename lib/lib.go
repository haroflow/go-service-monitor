package lib

import (
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// TestHTTPEndpoint sends a GET request to the specified address, with a
// default 10 second timeout, and expects a 200 OK response.
func TestHTTPEndpoint(address string) error {
	if !strings.HasPrefix(address, "http://") && !strings.HasPrefix(address, "https://") {
		return fmt.Errorf("address should begin with http:// or https://")
	}

	timeout := time.Duration(10 * time.Second)
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

// TestTCPEndpoint sends a TCP request to the specified address and port,
// with a default timeout of 10 seconds.
func TestTCPEndpoint(address string, port int16) error {
	timeout := time.Duration(10 * time.Second)
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

// TestDNSResponse does an IP lookup for the specified address.
func TestDNSResponse(address string) error {
	// FIXME we should be able to specify dns servers
	_, err := net.LookupIP(address)
	return err
}
