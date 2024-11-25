package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"

	"github.com/kaiiorg/receipt-processor/internal/api"
)

var (
	port = flag.Int("port", 8080, "Port to listen on for HTTP requests")
)

func main() {
	flag.Parse()
	go run()
	waitForInterrupt()
}

func run() {
	a := api.New()
	err := a.Run(fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		panic(err)
	}
}

func waitForInterrupt() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan
}
