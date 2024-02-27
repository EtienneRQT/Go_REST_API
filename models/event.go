package models

import (
	"time"

	"example.com/api/db"
)

// Event represents an event with ID, name, description, location, date/time, and user ID fields.
// The struct fields are validated as required.
type Event struct {
	ID          int64     `binding:"required"`
	Name        string    `binding:"required"`
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

// GetEventByID retrieves an Event by its ID from the database.
// It executes a SELECT query using the provided ID and scans the result
// into an Event struct. Returns the Event and any error from the query.
func GetEventByID(id int64) (Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events WHERE id =?"
	row := db.DB.QueryRow(query, id)
	var e Event
	err := row.Scan(
		&e.ID,
		&e.Name,
		&e.Description,
		&e.Location,
		&e.DateTime,
		&e.UserID)
	if err != nil {
		return Event{}, err
	}
	return e, nil
}

// GetAllEvents retrieves all events from the database.
// It executes a SELECT query to get all rows from the events table.
// The query result rows are scanned into Event structs and collected in a slice.
// Returns the slice of Event structs, or any error from the query.
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
