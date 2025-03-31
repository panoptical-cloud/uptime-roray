package main

import "github.com/nats-io/nats.go"

func NewNatsConn(url string) *nats.Conn {
	nc, err := nats.Connect(url)
	if err != nil {
		panic(err)
	}
	return nc
}
