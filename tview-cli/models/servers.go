package models

type ServerGroup struct {
	ID   int
	Name string
	Desc string
}

type ServerGroups []*ServerGroup

type Server struct {
	ID           string
	GroupId      int
	Ip           string
	Mac          string
	RegStatus    string
	OneTimeToken string
	Name         string
	Desc         string
	Fqdn         string
	AgentVersion string
	Os           string
	Arch         string
}

type Servers []*Server
