package service

import (
	"testing"
	"time"
)

func TestStoreEvents(t *testing.T) {
	// Create some sample events
	event1 := Event{
		OrderType:  "Purchase",
		SessionID:  "29827525-06c9-4b1e-9d9b-7c4584e82f56",
		Card:       "4433**1409",
		EventDate:  CustomTime{Time: time.Now()},
		WebsiteURL: "https://amazon.com",
	}
	event2 := Event{
		OrderType:  "CardVerify",
		SessionID:  "500cf308-e666-4639-aa9f-f6376015d1b4",
		Card:       "4433**1409",
		EventDate:  CustomTime{Time: time.Now()},
		WebsiteURL: "https://adidas.com",
	}

	// Store events
	err := StoreEvents([]Event{event1, event2})
	if err != nil {
		t.Errorf("error storing events: %v", err)
	}

	// Verify events are stored correctly
	storedEvents := GetStoredEvents()
	if len(storedEvents) != 2 {
		t.Errorf("expected 2 events to be stored, got %d", len(storedEvents))
	}

}
