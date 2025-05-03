package global

import (
	"pc-uptime/tview/models"

	"github.com/gdamore/tcell/v2"
)

func setCurrentKeyHandler(app *models.AppHolder) {
	// Set the current key handler for the application
	app.TviewApp.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlSpace:
			drawAppMainMenu(app)
		}
		return event
	})
}
