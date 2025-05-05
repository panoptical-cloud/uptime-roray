package main

import (
	"fmt"
	"pc-uptime/tview/api"

	"github.com/gdamore/tcell/v2"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"

	"github.com/rivo/tview"
)

// func DrawServersScreen(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func DrawServersScreen(app *tview.Application, focusOn tview.Primitive, f *tview.Flex, ssc *ServerScreenConfig) {
	LeftContent(app, focusOn, f, ssc)
	MiddleContent(app, f)
	RightContent(f)
}

// func LeftContent(f *tview.Flex, sgs []*ServerGroup, ss []*Server) {
func LeftContent(app *tview.Application, focusOn tview.Primitive, f *tview.Flex, ssc *ServerScreenConfig) {
	sgsList := ServerGroupList(ssc.ServerGroups)
	sgsList.SetTitle(" Server Groups (g) ").SetTitleAlign(tview.AlignLeft).SetBorder(true)
	ssList := ServerList(ssc)
	ssList.SetTitle(" Servers (s) ").SetTitleAlign(tview.AlignLeft).SetBorder(true)
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

func MiddleContent(app *tview.Application, f *tview.Flex) {
	metricsView := ServerMetricsView(app, f, "nats://107.155.65.50:4222", "agent.*.metrics.basic")
	smt := tview.NewTextView().SetText(" \nServer Metrics: \n")

	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" middle "), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(smt, 3, 1, false).
			AddItem(metricsView[0], 2, 1, false).
			AddItem(metricsView[1], 2, 1, false).
			AddItem(metricsView[2], 2, 1, false), 0, 1, false), 0, 2, false)
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
func ServerList(ssc *ServerScreenConfig) *tview.List {
	l := tview.NewList()
	for idx, s := range ssc.Servers {
		// l.AddItem(s.Name, s.IP, ToChar(idx), nil)
		l.AddItem(s.Name, s.IP, rune('1'+idx), func() {
			selSer := SelectedServer(ssc.Servers, idx)
			ssc.SelectedServer = selSer
		})
	}
	return l
}

func SelectedServerGroup(sgs []*ServerGroup, i int) *ServerGroup {
	if i < 0 || i >= len(sgs) {
		return nil
	}
	return sgs[i]
}

func SelectedServer(ss []*Server, i int) *Server {
	if i < 0 || i >= len(ss) {
		return nil
	}
	return ss[i]
}

// ServerMetricsView subscribes to a NATS subject for server metrics and displays them
func ServerMetricsView(app *tview.Application, f *tview.Flex, natsURL, subject string) []*tview.TextView {
	cpuTxt := tview.NewTextView().SetLabel("[yellow]CPU: ")
	diskTxt := tview.NewTextView().SetLabel("[yellow]Disk: ")
	memTxt := tview.NewTextView().SetLabel("[yellow]RAM: ")

	// Create buffered channel to avoid blocking on send
	smDC := make(chan *api.BaseStatsReply, 10)

	// Connect to NATS in a goroutine to avoid blocking UI
	go func() {
		// Connect to NATS server
		nc, err := nats.Connect(natsURL)
		if err != nil {
			app.QueueUpdateDraw(func() {
				memTxt.SetText(fmt.Sprintf("[red]Error connecting to NATS: %v[white]", err))
			})
			return
		}
		defer nc.Close()

		// Subscribe to the metrics subject - fix unused variable
		_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
			rcvData := &api.BaseStatsReply{}
			err = proto.Unmarshal(msg.Data, rcvData)
			if err != nil {
				return
			}
			smDC <- rcvData
		})

		if err != nil {
			app.QueueUpdateDraw(func() {
				memTxt.SetText(fmt.Sprintf("[red]Error subscribing to metrics: %v[white]", err))
			})
			return
		}
		// This is critical - without this, the goroutine could exit and close the NATS connection
		select {} // Block forever
	}()

	go func() {
		for {
			md := <-smDC

			// Update the TextView with new metrics data
			app.QueueUpdateDraw(func() {
				mem := fmt.Sprintf("%.2f %%", *md.Memory)
				cpu := fmt.Sprintf("%.2f %%", *md.Cpu)
				disk := fmt.Sprintf("%.2f %%", *md.Disk)
				memTxt.SetText(mem)
				cpuTxt.SetText(cpu)
				diskTxt.SetText(disk)
			})
		}
	}()
	return []*tview.TextView{memTxt, cpuTxt, diskTxt}
}
