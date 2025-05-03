package sg

import (
	"pc-uptime/tview/models"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func ViewServerGroupKeyHandlers(mf *tview.Flex, cf *tview.Flex, app *models.AppHolder, sgs []*ServerGroup) {
	mf.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'a':
			go func() {
				app.TviewApp.QueueUpdateDraw(func() {
					ht := app.GlobalElems.HandleFooterMsg("[esc] to go back")
					cf.RemoveItem(cf.GetItem(1))
					cf.AddItem(ht, 0, 3, false)

					form := DrawAddNewServerGroupForm(mf, cf, app)
					app.TviewApp.SetFocus(mf)
					mf.RemoveItem(mf.GetItem(1))
					mf.AddItem(form, 0, 8, true)
					AddNewServerGroupKeyHandler(mf, cf, app)
				})
			}()
		}
		return event
	})
}

func AddNewServerGroupKeyHandler(mf *tview.Flex, cf *tview.Flex, app *models.AppHolder) {
	mf.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyEscape:
			app.AppFlex.Clear()
			sgDC := make(chan []*ServerGroup)
			go GetAllServerGroupsSvc(sgDC)
			go DrawServerGroupScreen(app, sgDC)
		}
		return event
	})
}
