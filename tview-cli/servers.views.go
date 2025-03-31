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
	ssList := ServerList(ssc.Servers)
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
	metricsView := ServerMetricsView(app, "nats://localhost:4222", "agent.*.metrics.basic")
	// f.AddItem(tview.NewList().SetTitle("right").SetBorder(true), 0, 2, false)
	f.AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle(" middle "), 0, 1, false).
		AddItem(metricsView, 0, 1, false), 0, 2, false)
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

// ServerMetricsView subscribes to a NATS subject for server metrics and displays them
func ServerMetricsView(app *tview.Application, natsURL, subject string) *tview.TextView {
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWordWrap(true)

	textView.SetBorder(true).SetTitle(" Server Metrics ")
	textView.SetText("Connecting to NATS...")

	// Create buffered channel to avoid blocking on send
	smDC := make(chan *api.BaseStatsReply, 10)

	// Connect to NATS in a goroutine to avoid blocking UI
	go func() {
		// Connect to NATS server
		nc, err := nats.Connect(natsURL)
		if err != nil {
			app.QueueUpdateDraw(func() {
				textView.SetText(fmt.Sprintf("[red]Error connecting to NATS: %v[white]", err))
			})
			return
		}
		defer nc.Close()

		// Update the UI to show connection was successful
		app.QueueUpdateDraw(func() {
			textView.SetText("Connected to NATS. Waiting for metrics data...")
		})

		// fmt.Printf("Successfully connected to NATS at %s\n", natsURL)
		// fmt.Printf("Subscribing to subject: %s\n", subject)

		// Subscribe to the metrics subject - fix unused variable
		_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
			// fmt.Printf("Received message on subject: %s\n", msg.Subject)

			rcvData := &api.BaseStatsReply{}
			err = proto.Unmarshal(msg.Data, rcvData)
			if err != nil {
				// fmt.Printf("Error unmarshalling metrics data: %v\n", err)
				return
			}

			// fmt.Printf("Successfully unmarshalled message, sending to channel\n")
			smDC <- rcvData
		})

		if err != nil {
			app.QueueUpdateDraw(func() {
				textView.SetText(fmt.Sprintf("[red]Error subscribing to metrics: %v[white]", err))
			})
			return
		}

		// fmt.Printf("Successfully subscribed to: %s\n", subject)

		// This is critical - without this, the goroutine could exit and close the NATS connection
		select {} // Block forever
	}()

	go func() {
		for {
			// fmt.Printf("Waiting for metrics data on channel...\n")
			md := <-smDC
			// fmt.Printf("Received metrics data from channel: %+v\n", md)

			// Update the TextView with new metrics data
			app.QueueUpdateDraw(func() {
				s := fmt.Sprintf("%f", *md.Memory) // s == "123.456000"
				textView.SetText(s)
			})
		}
	}()

	return textView
}
