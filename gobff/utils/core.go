package utils

import (
	"net"
)

func GetIpFromHost(hostname string) (*string, error) {
	addrs, err := net.LookupIP(hostname)
	if err != nil {
		return nil, err
	}
	if len(addrs) == 0 {
		return nil, err
	}
	// Return the first IP address as string
	ret := addrs[0].String()
	return &ret, nil
}
