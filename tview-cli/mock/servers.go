package mock

import (
	t "pc-uptime/tview/models"
)

func ServerGroups() t.ServerGroups {
	var groups t.ServerGroups
	groups = append(groups, &t.ServerGroup{ID: 1, Name: "Group A", Desc: "Group 1 Description"})
	groups = append(groups, &t.ServerGroup{ID: 2, Name: "Group B", Desc: "Group 2 Description"})
	groups = append(groups, &t.ServerGroup{ID: 3, Name: "Group C", Desc: "Group 3 Description"})
	return groups
}

func ServersA() t.Servers {
	var servers t.Servers
	servers = append(servers, &t.Server{ID: "A1", GroupId: 1, Ip: "1.2.3.4", Mac: "00:11:22:33:44:55", RegStatus: "Registered", Name: "Server A1", Desc: "Server 1 Description", Fqdn: "server1.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "A2", GroupId: 1, Ip: "12.231.343.31", Mac: "00:11:22:33:44:56", RegStatus: "Registered", Name: "Server A2", Desc: "Server 2 Description", Fqdn: "server2.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "A3", GroupId: 1, Ip: "43.12.54.78", Mac: "00:11:22:33:44:57", RegStatus: "Registered", Name: "Server A3", Desc: "Server 3 Description", Fqdn: "server3.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "A4", GroupId: 1, Ip: "341.23.54.22", Mac: "00:11:22:33:44:58", RegStatus: "Registered", Name: "Server A4", Desc: "Server 4 Description", Fqdn: "server4.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	return servers
}

func ServersB() t.Servers {
	var servers t.Servers
	servers = append(servers, &t.Server{ID: "B1", GroupId: 2, Ip: "1.2.3.4", Mac: "00:11:22:33:44:55", RegStatus: "Registered", Name: "Server B1", Desc: "Server 1 Description", Fqdn: "server1.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "B2", GroupId: 2, Ip: "12.231.343.31", Mac: "00:11:22:33:44:56", RegStatus: "Registered", Name: "Server B2", Desc: "Server 2 Description", Fqdn: "server2.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "B3", GroupId: 2, Ip: "43.12.54.78", Mac: "00:11:22:33:44:57", RegStatus: "Registered", Name: "Server B3", Desc: "Server 3 Description", Fqdn: "server3.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	servers = append(servers, &t.Server{ID: "B4", GroupId: 2, Ip: "341.23.54.22", Mac: "00:11:22:33:44:58", RegStatus: "Registered", Name: "Server B4", Desc: "Server 4 Description", Fqdn: "server4.example.com", AgentVersion: "1.0.0", Os: "Linux", Arch: "x86_64"})
	return servers
}

func ServersByGroupId(groupId int) t.Servers {
	switch groupId {
	case 1:
		return ServersA()
	case 2:
		return ServersB()
	default:
		return nil
	}
}
