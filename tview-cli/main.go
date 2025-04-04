package main

import (
	"github.com/rivo/tview"
)

var (
	as *AppStateConfig
	af tview.Primitive
)

func main() {
	nc := NewNatsConn("nats://localhost:4222")
	as = &AppStateConfig{
		NC:            nc,
		CurrentScreen: WELCOME,
	}
	app := tview.NewApplication()
	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox().SetBorder(true).SetTitle(" Welcome "), 0, 1, false)
	sgDC := make(chan []*ServerGroup)
	sDC := make(chan []*Server)
	var sgs []*ServerGroup
	var ss []*Server
	go GetAllServerGroupsSvc(sgDC)
	sgs = <-sgDC
	go GetServersByGroupIdSvc(sgs[0].ID, sDC)
	ss = <-sDC
	as.ServerScreenConfig = &ServerScreenConfig{
		ServerGroups:        sgs,
		Servers:             ss,
		SelectedServerGroup: sgs[0],
		SelectedServer:      ss[0],
	}
	go func() {
		app.QueueUpdateDraw(func() {
			flex.Clear()
			DrawServersScreen(app, af, flex, as.ServerScreenConfig)
			app.SetFocus(flex)
		})
	}()
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
