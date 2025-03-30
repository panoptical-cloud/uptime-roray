package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// func DrawServersScreen(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func DrawServersScreen(f *tview.Flex, ssc *ServerScreenConfig) {
	LeftContent(f, ssc)
	MiddleContent(f)
	RightContent(f)
}

// func LeftContent(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func LeftContent(f *tview.Flex, ssc *ServerScreenConfig) {
	sgsList := ServerGroupList(ssc.ServerGroups)
	sgsList.SetTitle(" Server Groups ").SetBorder(true)
	ssList := ServerList(ssc.Servers)
	ssList.SetTitle(" Servers ").SetBorder(true).SetBackgroundColor(tcell.Color(0x000000))
	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(sgsList, 0, 1, true).
		AddItem(ssList, 0, 1, false), 0, 2, true)
}

func MiddleContent(f *tview.Flex) {
	f.AddItem(tview.NewList().SetTitle("middle").SetBorder(true), 0, 6, false)
}

func RightContent(f *tview.Flex) {
	f.AddItem(tview.NewList().SetTitle("right").SetBorder(true), 0, 2, false)
}

func ServerGroupList(sgs []*ServerGroup) *tview.List {
	l := tview.NewList()
	for idx, sg := range sgs {
		l.AddItem(sg.Name, sg.Desc, ToChar(idx), nil)
	}
	return l
}
func ServerList(ss []*Server) *tview.List {
	l := tview.NewList()
	for idx, s := range ss {
		l.AddItem(s.ID, s.Name, ToChar(idx), nil)
	}
	return l
}

func SelectedServerGroup(sgs []*ServerGroup, i int) *ServerGroup {
	if i < 0 || i >= len(sgs) {
		return nil
	}
	return sgs[i]
}
