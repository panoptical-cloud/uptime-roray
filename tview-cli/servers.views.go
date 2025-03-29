package main

import (
	"github.com/rivo/tview"
)

func DrawServersScreen(f *tview.Flex, sgs []*ServerGroup) {
	// func DrawServersScreen(f *tview.Flex, sgs []*ServerGroup,ss []*Server) {
	LeftContent(f, sgs)
	MiddleContent(f)
	RightContent(f)

}

func LeftContent(f *tview.Flex, sgs []*ServerGroup) {
	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewList().SetTitle(" Server Groups ").SetBorder(true), 0, 1, false).
		AddItem(tview.NewList().SetTitle("-"+sgs[0].Name+"-").SetBorder(true), 0, 1, false), 0, 2, false)
}

func MiddleContent(f *tview.Flex) {
	f.AddItem(tview.NewList().SetTitle("middle").SetBorder(true), 0, 6, false)
}

func RightContent(f *tview.Flex) {
	f.AddItem(tview.NewList().SetTitle("right").SetBorder(true), 0, 2, false)
}
