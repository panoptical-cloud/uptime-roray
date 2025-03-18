package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const tickTime = 3      // In seconds
const replyWaitTime = 3 // In seconds

func main() {
	app := &cli.App{
		Name:  "agent",
		Usage: "Collects and sends system metrics",
		Commands: []*cli.Command{
			{
				Name:  "register",
				Usage: "Register the agent with the central server",
				Action: func(c *cli.Context) error {
					err := RegisterCmd()
					if err != nil {
						log.Fatal(err)
						return err
					}
					return nil
				},
			},
			{
				Name:  "host",
				Usage: "Show host information",
				Action: func(c *cli.Context) error {
					err := HostInfoCmd()
					if err != nil {
						log.Fatal(err)
						return err
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			err := DefaultCmd()
			if err != nil {
				log.Fatal(err)
				return err
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
