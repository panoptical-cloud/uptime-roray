package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// func DrawServersScreen(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func DrawServersScreen(app *tview.Application, focusOn tview.Primitive, f *tview.Flex, ssc *ServerScreenConfig) {
	LeftContent(app, focusOn, f, ssc)
	MiddleContent(f)
	RightContent(f)
}

// func LeftContent(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func LeftContent(app *tview.Application, focusOn tview.Primitive, f *tview.Flex, ssc *ServerScreenConfig) {
	sgsList := ServerGroupList(ssc.ServerGroups)
	sgsList.SetTitle(" Server Groups (g) ").SetBorder(true)
	ssList := ServerList(ssc.Servers)
	ssList.SetTitle(" Servers (s) ").SetBorder(true)
	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(sgsList, 0, 1, true).
		AddItem(ssList, 0, 1, false), 0, 2, true)
	f.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Rune() {
		case 'g':
			app.SetFocus(sgsList)
		case 's':
			app.SetFocus(ssList)
		case 'q':
			app.Stop()
		}
		return event
	})
}

func MiddleContent(f *tview.Flex) {
	// f.AddItem(tview.NewList().SetTitle("right").SetBorder(true), 0, 2, false)
	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" middle "), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" middle "), 0, 1, false), 0, 2, false)
}

func RightContent(f *tview.Flex) {
	f.AddItem(tview.NewList().SetTitle("middle").SetBorder(true), 0, 6, false)
}

func ServerGroupList(sgs []*ServerGroup) *tview.List {
	l := tview.NewList()
	for idx, sg := range sgs {
		l.AddItem(sg.Desc, sg.Name, ToChar(idx), nil)
	}
	return l
}
func ServerList(ss []*Server) *tview.List {
	l := tview.NewList()
	for idx, s := range ss {
		l.AddItem(s.Name, s.IP, ToChar(idx), nil)
	}
	return l
}

func SelectedServerGroup(sgs []*ServerGroup, i int) *ServerGroup {
	if i < 0 || i >= len(sgs) {
		return nil
	}
	return sgs[i]
}
