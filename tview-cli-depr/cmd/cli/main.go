package main

import (
	"pc-uptime/tview/mock"
	v "pc-uptime/tview/views"

	"github.com/rivo/tview"
)

var as *v.AppState

func main() {

	// Initialize app state with defaults
	as = v.NewAppState()
	as.SetServers(mock.ServersA())

	appView := tview.NewApplication()

	// Create the main flex container (3 columns)
	flex := tview.NewFlex().SetDirection(tview.FlexColumn)

	// Create left column (3 rows)
	// leftColumn := tview.NewFlex().SetDirection(tview.FlexRow)
	// leftColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Left Top"), 0, 1, false)
	// leftColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Left Middle"), 0, 1, false)
	// leftColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Left Bottom"), 0, 1, false)

	// Create middle column (header, content, footer)
	middleColumn := tview.NewFlex().SetDirection(tview.FlexRow)
	middleColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Header"), 3, 1, false)
	middleColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Main Content"), 0, 3, false)
	middleColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Footer"), 3, 1, false)

	// Create right column (2 rows)
	rightColumn := tview.NewFlex().SetDirection(tview.FlexRow)
	rightColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Right Top"), 0, 1, false)
	rightColumn.AddItem(tview.NewBox().SetBorder(true).SetTitle("Right Bottom"), 0, 1, false)

	// Add columns to the main flex container
	flex.AddItem(v.LeftContent(as), 0, 1, true)
	flex.AddItem(middleColumn, 0, 6, false)
	flex.AddItem(rightColumn, 0, 1, false)

	if err := appView.SetRoot(flex, true).EnableMouse(true).EnablePaste(true).Run(); err != nil {
		panic(err)
	}

}
