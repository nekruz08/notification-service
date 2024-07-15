package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nekruz08/notification-service/pkg/handler"
	"github.com/nekruz08/notification-service/pkg/service"
)

func main() {
	// Start notification worker
	service.StartNotificationWorker()

	// HTTP server setup
	http.HandleFunc("/events", handler.EventHandler)
	fmt.Println("Server running on :8080...")

	// Start HTTP server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
