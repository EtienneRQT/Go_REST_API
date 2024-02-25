package models

import "time"

// Event represents an event with ID, name, description, location, date/time, and user ID fields.
// The struct fields are validated as required.
type Event struct {
	ID          int       `binding:"required"`
	Name        string    ` binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       `binding:"required"`
}

var events = []Event{}

// Save saves the Event e to the events slice.
func (e *Event) Save() {
	// save event to database
	events = append(events, *e)
}

// GetAllEvents returns all events.
func GetAllEvents() []Event {
	return events
}
