package main

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logrus.Error(err)
	}
	defer nc.Close()
	done := make(chan bool, 3)
	nc.Subscribe("foo", func(msg *nats.Msg) {
		fmt.Printf("Received a message: %s\n", msg.Data)
		done <- true
	})
	<-done
	<-done
}
