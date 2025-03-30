package main

import (
	"time"
)

func MockServerGroups() []*ServerGroup {
	// fmt.Println("MockServerGroups func called")
	time.Sleep(1 * time.Second)
	return []*ServerGroup{
		{1, "SG 1", "desc of SG 1"},
		{2, "SG 2", "desc of SG 2"},
		{3, "SG 3", "desc of SG 3"},
	}
}

func MockServersA() []*Server {
	time.Sleep(time.Second * 1)
	return []*Server{
		{ID: "1", GroupdId: 1, IP: "12.23.43.12", Name: "Server 1", Desc: "desc of server 1", Mac: "00:00:00:00:00:01", Os: "Linux", Arch: "x86_64"},
		{ID: "2", GroupdId: 1, IP: "42.31.54.64", Name: "Server 2", Desc: "desc of server 2", Mac: "00:00:00:00:00:02", Os: "Linux", Arch: "x86_64"},
		{ID: "3", GroupdId: 1, IP: "45.21.65.83", Name: "Server 3", Desc: "desc of server 3", Mac: "00:00:00:00:00:02", Os: "Linux", Arch: "x86_64"},
	}
}

func MockServersB() []*Server {
	time.Sleep(time.Second * 1)
	return []*Server{
		{ID: "4", GroupdId: 2, IP: "12.23.43.12", Name: "Server 4", Desc: "desc of server 1", Mac: "00:00:00:00:00:01", Os: "Linux", Arch: "x86_64"},
		{ID: "5", GroupdId: 2, IP: "42.31.54.64", Name: "Server 5", Desc: "desc of server 2", Mac: "00:00:00:00:00:02", Os: "Linux", Arch: "x86_64"},
		{ID: "6", GroupdId: 2, IP: "45.21.65.83", Name: "Server 6", Desc: "desc of server 3", Mac: "00:00:00:00:00:02", Os: "Linux", Arch: "x86_64"},
	}
}
