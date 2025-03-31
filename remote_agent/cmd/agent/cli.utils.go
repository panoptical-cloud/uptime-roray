package main

import (
	"fmt"
	"net"
	"pc-uptime/agent/api"
	"pc-uptime/agent/utils"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

func RegisterCmd() error {
	var registrationUrl string

	fmt.Print("Enter Registration URL: ")
	fmt.Scanln(&registrationUrl)
	fmt.Printf("Registering with server using : %s\n", registrationUrl)
	err := utils.RegisterWithServer(registrationUrl)
	if err != nil {
		return err
	}
	return nil
}

func HostInfoCmd() error {
	hostInfo, err := host.Info()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hostname: %s\n", hostInfo.Hostname)
	fmt.Printf("Host ID: %s\n", hostInfo.HostID)
	fmt.Printf("OS: %s\n", hostInfo.OS)
	fmt.Printf("Platform: %s %s\n", hostInfo.Platform, hostInfo.PlatformVersion)
	fmt.Printf("Kernel Version: %s\n", hostInfo.KernelVersion)
	fmt.Printf("Uptime: %d seconds\n", hostInfo.Uptime)

	// Format boot time into a readable date
	bootTime := time.Unix(int64(hostInfo.BootTime), 0)
	fmt.Printf("Boot Time: %s\n", bootTime.Format(time.RFC1123))

	// Show virtualization info if available
	if hostInfo.VirtualizationSystem != "" {
		fmt.Printf("Virtualization: %s (%s)\n",
			hostInfo.VirtualizationSystem,
			hostInfo.VirtualizationRole)
	}

	// Get and display IP addresses
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("Error getting network interfaces: %v\n", err)
	} else {
		fmt.Println("\nIP Addresses:")
		for _, i := range interfaces {
			addrs, err := i.Addrs()
			if err != nil {
				fmt.Printf("  Error getting addresses for interface %s: %v\n", i.Name, err)
				continue
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				// Filter out loopback addresses
				if !ip.IsLoopback() {
					fmt.Printf("  %s: %s\n", i.Name, ip.String())
				}
			}
		}
	}
	return nil
}

func DefaultCmd() error {
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
		cpuPercent, _ := cpu.Percent(0, false)
		diskUsage, _ := disk.Usage("/")

		stats := &api.BaseStatsReply{
			Cpu:    &cpuPercent[0],
			Memory: &v.UsedPercent,
			Disk:   &diskUsage.UsedPercent,
		}
		fmt.Printf("%#v\n", stats)

		msg, err := proto.Marshal(stats)
		if err != nil {
			log.Fatal(err)
		}

		// SEND
		// NATS - REQUEST & REPLY on "test.rpc" (THE PIPE)
		log.Printf("   Send request msg to subject 'agent.5eb1cf4a-da7d-426c-9642-eb15a411ccf6.metrics.basic'\n")
		_, err = nc.Request("agent.5eb1cf4a-da7d-426c-9642-eb15a411ccf6.metrics.basic", msg, replyWaitTime*time.Second)
		if err != nil {
			log.Println(err)
		}
	}
}
