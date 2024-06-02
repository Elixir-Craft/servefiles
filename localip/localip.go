package localip

import (
	"fmt"
	"net"
)

// Get returns a slice of local IP addresses as []net.IP
func Get() ([]net.IP, error) {
	// Get a list of all network addresses on the local system
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("error getting network addresses: %v", err)
	}

	var ips []net.IP

	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}

		// Skip loopback addresses and link-local addresses
		if ip == nil || ip.IsLoopback() || ip.IsLinkLocalUnicast() {
			continue
		}

		ips = append(ips, ip)
	}

	return ips, nil
}
