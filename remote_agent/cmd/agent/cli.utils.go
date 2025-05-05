package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"pc-uptime/agent/api"
	"pc-uptime/agent/utils"
	"strings"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
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
	fmt.Printf("Procs: %d\n", hostInfo.Procs)

	// Format boot time into a readable date
	bootTime := time.Unix(int64(hostInfo.BootTime), 0)
	fmt.Printf("Boot Time: %s\n", bootTime.Format(time.RFC1123))
	las, _ := load.Avg()
	fmt.Printf("Load Average: %v\n", las)

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

func ProcsInfoCmd(ctx context.Context) error {
	processes, err := process.ProcessesWithContext(ctx)
	if err != nil {
		log.Fatal(err)
		return err
	}
	tpp := "/usr/bin/sqlite3"
	for _, p := range processes {
		_cpp, _ := p.Exe()
		if tpp == _cpp {
			fmt.Printf("Running with Process ID: %d\n", p.Pid)
			since, err := p.CreateTimeWithContext(ctx)
			if err != nil {
				log.Fatal(err)
				return err
			}
			sinceTime := time.UnixMilli(since)
			fmt.Printf("Started at: %s\n", sinceTime.String())
		}
	}
	return nil
}

func DefaultCmda() error {
	// Default action: run the agent and send metrics
	// Load configuration (e.g., server URL, authentication token)
	// from a file or environment variables.
	// For now, assume NATS is running locally.

	nc, err := nats.Connect("nats://107.155.65.50:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	log.Println("Connected to nats://107.155.65.50:4222")

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

func DefaultCmd() error {

	ret, err := getInitConf()
	if err != nil {
		log.Fatal(err)
		return err
	}
	nc, err := nats.Connect(ret.NatsUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer nc.Close()
	log.Printf("Connected to %s: \n", ret.NatsUrl)

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
		subj := fmt.Sprintf("agent.%s.metrics.basic", ret.NatsSubject)
		log.Printf("   Send request msg to subject '%s'\n", subj)
		err = nc.Publish(subj, msg)
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
}

func getInitConf() (*utils.InitConf, error) {
	ret := &utils.InitConf{}
	configFile := ".config/panmon/agent.toml"
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %w", err)
	}
	agentHomePath := filepath.Join(homeDir, configFile)
	fb, err := os.ReadFile(agentHomePath)
	if err != nil {
		return nil, err
	}
	f := string(fb)
	ks := strings.Split(f, "\n")
	for _, k := range ks {
		_s := strings.Split(k, " = ")
		if _s[0] == "nats_url" {
			ret.NatsUrl = _s[1]
		}
		if _s[0] == "nats_subject" {
			ret.NatsSubject = _s[1]
		}
	}
	println(ret.NatsUrl)
	println(ret.NatsSubject)
	return ret, nil
}
