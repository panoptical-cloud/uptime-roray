package views

import (
	"pc-uptime/tview/mock"
	m "pc-uptime/tview/models"
)

func (as *AppState) ToggleSelectedServerGrp(sg *m.ServerGroup) {
	as.SetSelectedServerGrp(sg)
	as.SetServers(mock.ServersByGroupId(sg.ID))
}
