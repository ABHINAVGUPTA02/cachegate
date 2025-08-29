package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var server *http.Server

func startServer(addr string) {
	if server != nil {
		log.Println("Server already running")
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello! Current time: %s\n", time.Now())
	})

	server = &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Printf("Starting server on %s\n", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}

func stopServer() {
	if server == nil {
		return
	}
	log.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Shutdown error: %v", err)
	}
	server = nil
}

func restartServer(addr string) {
	stopServer()
	startServer(addr)
}
