package main

import (
	"os"
	"os/signal"
	"syscall"
)

func signal_demo() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT)
	<-interrupt
	<-interrupt
	<-interrupt
}
