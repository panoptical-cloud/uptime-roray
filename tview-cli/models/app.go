package models

import (
	"github.com/gdamore/tcell/v2"
	"github.com/nats-io/nats.go"
	"github.com/rivo/tview"
)

type AppHolder struct {
	TviewApp    *tview.Application
	AppFlex     *tview.Flex
	GlobalElems GlobalHandler
	NC          *nats.Conn
}

type GlobalHandler interface {
	HandleFooterMsg(txt string) tview.Primitive
}
type PanmonTabOpts struct {
	hc *tcell.Color
	dc *tcell.Color
}

func NewPanmonTable[T any](title string, headers []string, data []*T, mapper func(item T) []string, opts *PanmonTabOpts) *tview.Table {
	//Setting General table attributes and defaults
	hc := tcell.ColorTeal
	dc := tcell.ColorWhiteSmoke
	t := tview.NewTable().SetBorders(true)
	t.SetBorder(true)
	t.SetTitle(title)
	t.SetTitleAlign(tview.AlignLeft)

	// Setting header & data cell colors thru opts
	if opts != nil {
		if opts.hc != nil {
			hc = *opts.hc
		}
		if opts.dc != nil {
			dc = *opts.dc
		}
	}

	// Setting header cells i.e. row=0
	for i, h := range headers {
		t.SetCell(0, i, tview.NewTableCell(h).SetTextColor(hc).SetAlign(tview.AlignCenter))
	}

	// Setting data cells
	for ri, item := range data {
		for ci, v := range mapper(*item) {
			t.SetCell(ri+1, ci, tview.NewTableCell(v).SetTextColor(dc).SetAlign(tview.AlignLeft))
		}
	}
	return t
}
