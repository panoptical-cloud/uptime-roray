package main

import (
	"fmt"
	"pc-uptime/agent/api"

	"google.golang.org/protobuf/proto"

	"time"

	"github.com/nats-io/nats.go"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
	log "github.com/sirupsen/logrus"
)

const tickTime = 3      // In seconds
const replyWaitTime = 3 // In seconds

func main() {
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	log.Println("Connected to nats://localhost:4222")

	for c := time.Tick(time.Duration(tickTime) * time.Second); ; <-c {

		v, _ := mem.VirtualMemory()
		c, _ := cpu.Percent(1, false)
		d, _ := disk.Usage("/")
		// almost every return value is a struct
		fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

		// convert to JSON. String() is also implemented
		// fmt.Println(v)

		stats := &api.BaseStatsReply{
			Hostname: "test",
			Cpu:      int32(c[0]),
			Memory:   int32(v.UsedPercent),
			Disk:     int32(d.UsedPercent),
		}

		// MARSHAL
		msg, err := proto.Marshal(stats)
		if err != nil {
			log.Fatal(err)
		}

		// SEND
		// NATS - REQUEST & REPLY on "test.rpc" (THE PIPE)
		log.Printf("   Send request msg to subject 'test.rpc'\n")
		_, err = nc.Request("test.rpc", msg, replyWaitTime*time.Second)
		if err != nil {
			log.Println(err)
		}

	}
}
