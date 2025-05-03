package sg

import (
	"errors"
	"fmt"
	"log"
	"pc-uptime/tview/api"
	m "pc-uptime/tview/models"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/rivo/tview"
	"google.golang.org/protobuf/proto"
)

func DrawAddNewServerGroupForm(mf *tview.Flex, cf *tview.Flex, app *m.AppHolder) *tview.Form {
	var newGN *string
	var newGD *string
	form := tview.NewForm().
	AddInputField("Group Name", "", 20, nil, func(text string) {
		newGN = &text
	}).
	AddTextArea("Group Description", "", 40,0,0, func(text string) {
		newGD = &text
	}).
	AddButton("Save", func() {
		err:=AddNewServerGroup(*newGN, *newGD)
		if err != nil {
			panic(err)
		}
		app.AppFlex.Clear()
		sgDC:= make(chan []*ServerGroup)
		go GetAllServerGroupsSvc(sgDC)
		go DrawServerGroupScreen(app, sgDC)
	})
	form.SetBorder(true).SetTitle(" Add New Server Group ").SetTitleAlign(tview.AlignCenter)
	return form
}

func drawServersBaseMetricsOverviewTable(app *m.AppHolder, f *tview.Flex, sg *ServerGroup, sDC chan []*Server) {
	// Implementation of the function to draw the server base metrics overview table
	// This function will be called when a server group is selected from the list
	// It will use the provided channel to get the servers for the selected group
	if f.GetItemCount() > 0 {
		f.RemoveItem(f.GetItem(0))
		f.Clear()
	}

	servers := <-sDC
	bsDC := make([]chan *api.BaseStatsReply, len(servers))

	headers := []string{"Name", "FQDN", "IP", "Uptime", "Agent", "CPU", "RAM", "Disk"}
	t := m.NewPanmonTable("[Group: "+sg.Name+"]", headers, servers, func(s Server) []string {
		return []string{s.Name, s.Fqdn, s.IP, "Up since 6D 2H 31M 21S", "Online | v: 1.2.3", "", "", ""}
	}, nil)

	for idx := range servers {
		bsDC[idx] = make(chan *api.BaseStatsReply)
		go func() {
			_, err := app.NC.Subscribe("agent."+"1.2.1.2"+".metrics.basic", func(msg *nats.Msg) {
				data := &api.BaseStatsReply{}
				err := proto.Unmarshal(msg.Data, data)
				if err != nil {
					log.Fatal(err)
				}
				bsDC[idx] <- data
			})
			if err != nil {
				log.Fatal(errors.Unwrap(err))
			}
			select {}
		}()
	}

	for idx, ch := range bsDC {
		go func(channel chan *api.BaseStatsReply, index int) {
			for c := time.Tick(time.Duration(1) * time.Second); ; <-c {
				rcvData := <-ch
				app.TviewApp.QueueUpdateDraw(func() {
					cpu_str := fmt.Sprintf("%.2f", *rcvData.Cpu)
					t.GetCell(idx+1, 5).SetText(cpu_str)
					mem_str := fmt.Sprintf("%.2f", *rcvData.Memory)
					t.GetCell(idx+1, 6).SetText(mem_str)
					disk_str := fmt.Sprintf("%.2f", *rcvData.Disk)
					t.GetCell(idx+1, 7).SetText(disk_str)
				})
			}
		}(ch, idx)
	}
	f.AddItem(t, 0, 1, false)
	t.SetSelectable(true, false)
	t.SetBorderPadding(1, 0, 1, 0)
	t.SetSelectedFunc(func(r, c int) {
		go func() {
			app.TviewApp.QueueUpdateDraw(func() {
				drawServerSvcStatusOverviewTable(app, f)
			})
		}()
	})
}

func drawServerSvcStatusOverviewTable(app *m.AppHolder, f *tview.Flex) {
	if f.GetItemCount() == 2 {
		f.RemoveItem(f.GetItem(1))
	}

	md := make([]*ServerSvcStats, 3)
	md[0] = &ServerSvcStats{
		Path:    "/opt/app/nginx/web",
		Port:    8080,
		Running: true,
		Since:   "01-Apr-2025 6:30 UTC",
	}
	md[1] = &ServerSvcStats{
		Path:    "/opt/app/apache/web",
		Port:    8181,
		Running: true,
		Since:   "01-Apr-2025 8:30 UTC",
	}
	md[2] = &ServerSvcStats{
		Path:    "/opt/app/etl/order_sync_batch",
		Port:    -1,
		Running: true,
		Since:   "02-Apr-2025 18:30 UTC",
	}
	headers := []string{"Path", "Port", "Running", "Since"}
	t := m.NewPanmonTable("[Service Status: abc.com (1.2.1.2)] ", headers, md, func(s ServerSvcStats) []string {
		portStr := strconv.Itoa(int(s.Port))
		isRunning := "NO"
		if s.Running {
			isRunning = "YES"
		}
		return []string{s.Path, portStr, isRunning, s.Since}
	}, nil)
	f.AddItem(t, 0, 1, false)
}
