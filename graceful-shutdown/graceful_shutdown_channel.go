package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Println("Server started")
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// os.Interrupt		-> control+c  or kill -SIGINT
	// syscall.SIGTERM	-> kill <pid>

	log.Println("wait for signal")

	<-stop
	log.Println("server stopped")
}
