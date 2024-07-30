package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e *Event) Save() {
	// later: store in the DB
	events = append(events, *e)
}

func GetAllEvents() []Event {
	// later: DB request
	return events
}
