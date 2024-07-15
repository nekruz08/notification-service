package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nekruz08/notification-service/pkg/service"

)

func TestEventHandler(t *testing.T) {
	// Create a sample JSON payload
	payload := `[{
		"orderType": "Purchase",
		"sessionId": "29827525-06c9-4b1e-9d9b-7c4584e82f56",
		"card": "4433**1409",
		"eventDate": "2023-01-04 13:44:52.835626 +00:00",
		"websiteUrl": "https://amazon.com"
	}]`

	// Create a request with the sample payload
	req, err := http.NewRequest("POST", "/events", bytes.NewBuffer([]byte(payload)))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to capture the handler's response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(EventHandler)

	// Serve the HTTP request to the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Verify the response body
	expected := "Events received successfully!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}

	// Verify that events were stored correctly
	storedEvents := service.GetStoredEvents()
	if len(storedEvents) != 1 {
		t.Errorf("expected 1 event to be stored, got %d", len(storedEvents))
	}

}
