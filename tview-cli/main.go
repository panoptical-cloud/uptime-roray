package main

import (
	g "pc-uptime/tview/global"
	m "pc-uptime/tview/models"
	sg "pc-uptime/tview/sg"
)

var App *m.AppHolder

func main() {
	App = g.NewApp("nats://107.155.65.50:4222")
	defer App.NC.Close()
	sgDC := make(chan []*sg.ServerGroup)
	go sg.GetAllServerGroupsSvc(sgDC)
	go sg.DrawServerGroupScreen(App, sgDC)
	if err := App.TviewApp.SetRoot(App.AppFlex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}
}
