package service

import (
	"fmt"
	"sync"
	"time"
)

// Event represents the JSON structure of an event
type Event struct {
	OrderType  string    `json:"orderType"`
	SessionID  string    `json:"sessionId"`
	Card       string    `json:"card"`
	EventDate  CustomTime `json:"eventDate"`
	WebsiteURL string    `json:"websiteUrl"`
}

// CustomTime is a custom type that implements encoding/json.Unmarshaler
// to parse a specific date format.
type CustomTime struct {
	time.Time
}

const customTimeLayout = "2006-01-02 15:04:05.999999999 -07:00"

// UnmarshalJSON parses a JSON date string in a specific format into time.Time.
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	// Removes quotes from the JSON string
	s := string(b)
	s = s[1 : len(s)-1]
	t, err := time.Parse(customTimeLayout, s)
	if err != nil {
		return err
	}
	ct.Time = t
	return nil
}

var (
	events   []Event
	eventsMu sync.Mutex
)

// StoreEvents stores events in memory
func StoreEvents(receivedEvents []Event) error {
	eventsMu.Lock()
	defer eventsMu.Unlock()

	events = append(events, receivedEvents...)
	return nil
}

// GetStoredEvents retrieves the currently stored events.
func GetStoredEvents() []Event {
	eventsMu.Lock()
	defer eventsMu.Unlock()

	// Create a copy of events to avoid race conditions
	// when accessing it concurrently.
	copiedEvents := make([]Event, len(events))
	copy(copiedEvents, events)

	return copiedEvents
}

// notificationWorker runs a worker to process events and notify clients (mocked by printing to terminal)
func StartNotificationWorker() {
	go notificationWorker()
}

func notificationWorker() {
	for {
		eventsMu.Lock()
		for _, event := range events {
			fmt.Printf("Notification: Event %+v processed.\n", event)
		}
		events = nil // Clear events after processing
		eventsMu.Unlock()
		time.Sleep(5 * time.Second) // We coudl adjust sleep time as needed
	}
}
