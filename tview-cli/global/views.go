package global

import (
	m "pc-uptime/tview/models"
	sg "pc-uptime/tview/sg"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func drawAppFooterText(txt string) *tview.TextView {
	txt = " " + txt + " [(ctrl+space) to open main menu] [(ctrl+c) to exit] "
	ft := tview.NewTextView().SetText(txt)
	ft.SetBackgroundColor(tcell.ColorDarkSlateGray)
	return ft
}

func drawAppMainMenu(app *m.AppHolder) {
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	list := tview.NewList().
		AddItem("Servers", "Monitor server vitals, processes, stream logs etc.", '1', func() {
			sgDC := make(chan []*sg.ServerGroup)
			go sg.GetAllServerGroupsSvc(sgDC)
			app.AppFlex.Clear()
			go sg.DrawServerGroupScreen(app, sgDC)
		}).
		AddItem("URLs", "Monitor http(s) endpoints or URLs", '2', nil).
		AddItem("Databases", "Monitor databases", '3', nil).
		AddItem("Kubernetes", "Monitor kubernetes clusters", '4', nil).
		AddItem("Quit", "Quit the app", 'q', func() {
			app.TviewApp.Stop()
		})
	list.SetBorder(true)
	list.SetTitle(" Main Menu ")
	pages := tview.NewPages().
		AddPage("modal", modal(list, 40, 20), true, true)
	app.AppFlex.Clear()
	app.AppFlex.AddItem(pages, 0, 1, true)
	app.TviewApp.SetFocus(app.AppFlex)
}
