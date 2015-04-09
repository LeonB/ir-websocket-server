package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Test")

	// http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws/session", newSessionHandler())
	// http.HandleFunc("/ws/telemetry", serveWs)
	// http.HandleFunc("/ws/remote", serveWs)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
