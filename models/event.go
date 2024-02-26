package models

import (
	"time"

	"example.com/api/db"
)

// Event represents an event with ID, name, description, location, date/time, and user ID fields.
// The struct fields are validated as required.
type Event struct {
	ID          int64     `binding:"required"`
	Name        string    ` binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int       `binding:"required"`
}

// Save inserts the Event into the database.
// It prepares an INSERT statement using the Event fields, executes it,
// retrieves the inserted ID, and sets it on the Event.
// Returns any errors from preparing, executing the statement,
// or retrieving the inserted ID.
func (e *Event) Save() error {
	query := "INSERT INTO events (name, description, location, dateTime, user_id) VALUES (?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	if err != nil {
		return err
	}

	return nil
}

// GetAllEvents returns all events.
func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var events []Event
	for rows.Next() {
		var e Event
		err = rows.Scan(
			&e.ID,
			&e.Name,
			&e.Description,
			&e.Location,
			&e.DateTime,
			&e.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	return events, nil
}
