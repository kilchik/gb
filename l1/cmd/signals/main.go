package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Print("start")

	sigintChan := make(chan os.Signal)
	signal.Notify(sigintChan, syscall.SIGINT)

	sigtermChan := make(chan os.Signal)
	signal.Notify(sigtermChan, syscall.SIGTERM)

	select {
	case <-sigintChan:
		log.Println("got sigint")

	case <-sigtermChan:
		log.Println("got sigterm")
	}

	log.Print("end")
}
