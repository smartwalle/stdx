package netx

import (
	"net"
)

func HardwareAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		var addr = iface.HardwareAddr.String()
		if len(addr) != 0 {
			return addr, nil
		}
	}
	return "", nil
}

func HardwareAddrs() ([]string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var addrs = make([]string, 0, len(interfaces))
	for _, iface := range interfaces {
		var addr = iface.HardwareAddr.String()
		if len(addr) != 0 {
			addrs = append(addrs, addr)
		}
	}
	return addrs, nil
}
