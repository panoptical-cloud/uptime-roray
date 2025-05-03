package sg

import (
	api "pc-uptime/tview/api"
)

type ServerGroup struct {
	ID   int
	Name string
	Desc string
}

type Server struct {
	ID      string
	GroupId int
	IP      string
	Name    string
	Desc    string
	Mac     string
	Os      string
	Arch    string
	Fqdn    string
}

type ServerBaseStats struct {
	IP    string
	Stats *api.BaseStatsReply
}

type ServerSvcStats struct {
	IP      string
	Port    int32
	Path    string
	Running bool
	Since   string
}
