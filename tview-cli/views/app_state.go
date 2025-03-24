package views

import (
	m "pc-uptime/tview/models"
)

type AppState struct {
	servers           []*m.Server
	serverGroups      []*m.ServerGroup
	selectedServerGrp *m.SelectedServerGrp
	selectedServer    *m.SelectedServer
}

func NewAppState() *AppState {
	return &AppState{}
}

func (as *AppState) GetSelectedServerGrp() *m.ServerGroup {
	if as.selectedServerGrp == nil {
		return nil
	}
	return (*as.selectedServerGrp)
}

func (as *AppState) SetSelectedServerGrp(selectedServerGrp *m.ServerGroup) {
	ssg := m.SelectedServerGrp(selectedServerGrp)
	as.selectedServerGrp = &ssg
}

func (as *AppState) GetSelectedServer() *m.Server {
	if as.selectedServer == nil {
		return nil
	}
	return (*as.selectedServer)
}

func (as *AppState) SetSelectedServer(selectedServer *m.Server) {
	ss := m.SelectedServer(selectedServer)
	as.selectedServer = &ss
}

func (as *AppState) GetServerGroups() []*m.ServerGroup {
	return as.serverGroups
}

func (as *AppState) SetServerGroups(serverGroups []*m.ServerGroup) {
	as.serverGroups = serverGroups
}

func (as *AppState) GetServers() []*m.Server {
	return as.servers
}

func (as *AppState) SetServers(servers []*m.Server) {
	as.servers = servers
}
