package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

var (
	port  int
	delay int
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen on")
	flag.IntVar(&delay, "delay", 0, "Delay in seconds")
	flag.Parse()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok\n"))
	})

	http.HandleFunc("/delay", func(w http.ResponseWriter, r *http.Request) {
		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Second)
		}
		w.Write([]byte("ok\n"))
	})

	// Start the server
	fmt.Printf("Listening on :%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
