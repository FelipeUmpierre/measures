package repository

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	lk "github.com/ulule/loukoum"
)

type (
	// EventRepo holds the database connection
	EventRepo struct {
		db *sqlx.DB
	}

	// Event interface for the event
	Event interface {
		Handle() string
		AggregateID() string
		Payload() json.RawMessage
		CreatedAt() time.Time
	}
)

// NewEventRepository creates a new repository
func NewEventRepository(db *sqlx.DB) *EventRepo {
	return &EventRepo{db}
}

// Emit source the event
func (m *EventRepo) Emit(e Event) error {
	query, args := lk.Insert(`events`).Set(
		lk.Pair(`handle`, e.Handle()),
		lk.Pair(`aggregate_id`, e.AggregateID()),
		lk.Pair(`payload`, string(e.Payload())),
		lk.Pair(`created_at`, e.CreatedAt()),
	).Prepare()

	stmt, err := m.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args)
	return err
}

// EventsFrom search for all events from a specific aggregate_id
// func (m *MeasuresRepo) EventsFrom(ID string) (*[]event.Event, error) {
// 	query, args := lk.Select(`id`, `handle`, `aggregate_id`, `payload`).
// 		From(`measures`).
// 		Where(lk.Condition(`aggregate_id`).Equal(ID)).
// 		Prepare()

// 	stmt, err := m.db.PrepareNamed(query)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer stmt.Close()

// 	events := new([]event.Event)
// 	err = stmt.Select(events, args)

// 	return measures, err
// }
