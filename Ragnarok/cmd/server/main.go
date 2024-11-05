package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FOXHOUND0x/ragnarok/internal/api"
)

func main() {
	port := 8080
	fmt.Printf("Starting API server on port %d\n", port)

	// Initialize the API server
	http.HandleFunc("/containers", api.HandleContainers)

	// Start the server
	addr := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
