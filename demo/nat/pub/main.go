package main

import (
	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Error(err)
	}
	defer nc.Close()
	nc.Publish("foo", []byte("hello world"))
}
