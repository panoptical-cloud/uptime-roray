package main

import (
	"fmt"
	"os"
	"pc-uptime/agent/api"
	"pc-uptime/agent/utils"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/protobuf/proto"
)

const tickTime = 3      // In seconds
const replyWaitTime = 3 // In seconds

func main() {
	app := &cli.App{
		Name:  "agent",
		Usage: "Collects and sends system metrics",
		Commands: []*cli.Command{
			{
				Name:    "register",
				Aliases: []string{"r"},
				Usage:   "Register the agent with the central server",
				Action: func(c *cli.Context) error {
					var registrationUrl string

					fmt.Print("Enter Registration URL: ")
					fmt.Scanln(&registrationUrl)
					fmt.Printf("Registering with server using : %s\n", registrationUrl)
					err := utils.RegisterWithServer(registrationUrl)
					if err != nil {
						log.Fatal(err)
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			// Default action: run the agent and send metrics
			// Load configuration (e.g., server URL, authentication token)
			// from a file or environment variables.
			// For now, assume NATS is running locally.

			nc, err := nats.Connect("nats://localhost:4222")
			if err != nil {
				log.Fatal(err)
			}
			defer nc.Close()
			log.Println("Connected to nats://localhost:4222")

			for c := time.Tick(time.Duration(tickTime) * time.Second); ; <-c {

				v, _ := mem.VirtualMemory()
				cpuPercent, _ := cpu.Percent(1, false)
				diskUsage, _ := disk.Usage("/")
				// almost every return value is a struct
				fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

				// convert to JSON. String() is also implemented
				// fmt.Println(v)

				stats := &api.BaseStatsReply{
					Hostname: "test",
					Cpu:      int32(cpuPercent[0]),
					Memory:   int32(v.UsedPercent),
					Disk:     int32(diskUsage.UsedPercent),
				}

				// MARSHAL
				msg, err := proto.Marshal(stats)
				if err != nil {
					log.Fatal(err)
				}

				// SEND
				// NATS - REQUEST & REPLY on "test.rpc" (THE PIPE)
				log.Printf("   Send request msg to subject 'agent.localhost.metrics.basic'\n")
				_, err = nc.Request("agent.localhost.metrics.basic", msg, replyWaitTime*time.Second)
				if err != nil {
					log.Println(err)
				}

			}
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
