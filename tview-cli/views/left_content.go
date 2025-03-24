package views

import (
	mock "pc-uptime/tview/mock"
	"pc-uptime/tview/models"
	u "pc-uptime/tview/utils"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func LeftContent(as *AppState) *tview.Flex {
	// Create left column (3 rows)
	menu := tview.NewFlex().SetDirection(tview.FlexRow)
	menu.AddItem(leftTopMenu(as), 0, 2, true)
	menu.AddItem(leftMiddleMenu(), 0, 2, false)
	menu.AddItem(leftBottomMenu(), 0, 1, false)
	return menu
}

func leftTopMenu(as *AppState) tview.Primitive {
	// v := tview.NewBox().SetBorder(true).SetTitle(" Server Groups (1) ").SetBorderColor(tcell.ColorOrange).SetTitleColor(tcell.ColorDarkOrange)
	v := tview.NewList().ShowSecondaryText(false)
	v.SetTitle(" Server Groups (1) ").SetBorder(true).SetTitleColor(tcell.ColorDarkOrange).SetBorderColor(tcell.ColorOrange)
	sg := mock.ServerGroups()
	for idx, g := range sg {
		v.AddItem(g.Name, g.Desc, u.ToChar(idx+1), func() {
			as.ToggleSelectedServerGrp(g)
		})
	}
	return v
}

func leftMiddleMenu(incServers models.Servers) tview.Primitive {
	v := tview.NewList().ShowSecondaryText(false)
	v.SetTitle(" Servers (2) ").SetBorder(true).SetTitleColor(tcell.ColorAzure).SetBorderColor(tcell.ColorBlue)
	for idx, sv := range incServers {
		v.AddItem(sv.Name, sv.Desc, u.ToChar(idx+1), nil)
	}
	return v
}

func leftBottomMenu() tview.Primitive {
	v := tview.NewBox().SetBorder(true).SetTitle(" Help (3) ")
	return v
}
