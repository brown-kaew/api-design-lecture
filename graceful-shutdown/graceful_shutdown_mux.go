package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("Server start...")
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`OK`))
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`Hi`))
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// start server in separate process
	go func() {
		log.Fatal(srv.ListenAndServe())
	}()
	fmt.Println("server starting at :8080")

	// listening signal
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	fmt.Println("shutting down...")
	// stop recieving request and finish remain process
	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Println("shutdown err:", err)
	}
	fmt.Println("bye bye")
}
