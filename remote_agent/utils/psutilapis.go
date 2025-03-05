package utils

import (
	"fmt"
	"log"
	"net"

	"github.com/shirou/gopsutil/host"
)

type PsUtilsProvider interface {
	GetSelfIp() (*string, error)
	GetSelfMachineId() (*string, error)
}

type PsUtils struct{}

func (*PsUtils) GetSelfIp() (*string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	for _, addr := range addrs {
		ip, _, err := net.ParseCIDR(addr.String())
		if err != nil {
			return nil, err
		}
		if ip.IsGlobalUnicast() {
			_ip := ip.String()
			log.Printf("Found global unicast IP: %s", _ip)
			return &_ip, nil
		}
	}
	return nil, fmt.Errorf("no global unicast IP found")
}

func (*PsUtils) GetSelfMachineId() (*string, error) {
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &hostInfo.HostID, nil
}
