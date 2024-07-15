package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/nekruz08/notification-service/pkg/service"
)

// EventHandler handles POST requests to /events
func EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request body into slice of Event structs
	var receivedEvents []service.Event
	if err := json.NewDecoder(r.Body).Decode(&receivedEvents); err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Pass events to service for processing and storage
	if err := service.StoreEvents(receivedEvents); err != nil {
		http.Error(w, "Failed to store events", http.StatusInternalServerError)
		return
	}

	// Mock notification by printing the received events
	for _, event := range receivedEvents {
		fmt.Printf("Received event: %+v\n", event)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Events received successfully!"))
}
