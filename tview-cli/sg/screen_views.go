package sg

import (
	m "pc-uptime/tview/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func DrawServerGroupScreen(app *m.AppHolder, sgDC chan []*ServerGroup) {
	sgs := <-sgDC
	mf := tview.NewFlex()
	c := tview.NewFlex()
	nav := tview.NewList()
	nav.SetBorder(true).SetBorderColor(tcell.ColorTeal).SetTitle(" [Server Groups] ").SetTitleAlign(tview.AlignLeft)
	nav.SetChangedFunc(func(index int, mainText string, secondaryText string, shortcut rune) {
		sDC := make(chan []*Server)
		go GetServersByGroupIdSvc(sgs[index].ID, sDC)
		go func() {
			app.TviewApp.QueueUpdateDraw(func() {
				drawServersBaseMetricsOverviewTable(app, c, sgs[index], sDC)
			})
		}()
	})
	for i, sg := range sgs {
		r := rune('1' + i)
		nav.AddItem(sg.Name, sg.Desc, r, func() {
			//TODO: cb func when a server group is selected, switch to servers view
		})
	}
	mf.AddItem(nav, 0, 2, true).AddItem(c, 0, 8, false)
	ht := app.GlobalElems.HandleFooterMsg("[a] to add new server group")
	cf := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(mf, 0, 200, true).
		AddItem(ht, 0, 3, false)
	app.AppFlex.AddItem(cf, 0, 1, true)
	app.TviewApp.SetFocus(mf)
	ViewServerGroupKeyHandlers(mf, cf, app, sgs)
	close(sgDC)
}
