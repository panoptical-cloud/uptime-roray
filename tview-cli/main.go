package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox().SetBorder(true).SetTitle(" Welcomea "), 0, 1, false)
	sgDC := make(chan []*ServerGroup)
	go GetAllServerGroupsSvc(sgDC)
	go func() {
		sgs := <-sgDC
		app.QueueUpdateDraw(func() {
			flex.Clear()
			// fmt.Println("sgs", sgs)
			DrawServersScreen(flex, sgs)
			// flex.AddItem(tview.NewBox().SetBorder(true).SetTitle(sgs[0].Name), 0, 1, false)
		})
	}()
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}

}
