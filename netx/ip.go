package netx

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

var (
	ErrFailedToObtainIP = errors.New("failed to obtain IP address")
)

func IPv4ToInt(IPv4Addr string) (uint32, error) {
	ip := net.ParseIP(IPv4Addr)
	if ip == nil {
		return 0, fmt.Errorf("invalid IP address: %s", IPv4Addr)
	}

	ipv4 := ip.To4()
	if ipv4 == nil {
		return 0, fmt.Errorf("not an IPv4 address: %s", IPv4Addr)
	}

	return binary.BigEndian.Uint32(ipv4), nil
}

func ExternalIPv6() (string, error) {
	rsp, err := http.Get("https://ipv6.myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func ExternalIPv4() (string, error) {
	rsp, err := http.Get("https://ipv4.myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer rsp.Body.Close()

	content, err := io.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(content)), nil
}

func InternalIPv4() (string, error) {
	ips, err := InternalIPv4s()
	if err != nil {
		return "", err
	}
	if len(ips) == 0 {
		return "", ErrFailedToObtainIP
	}
	return ips[0], nil
}

func InternalIPv4s() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	var ips = make([]string, 0, len(addrs))
	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips, nil
}
