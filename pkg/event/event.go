package event

import (
	"encoding/json"
	"time"
)

type (
	// Event default struct
	Event struct {
		handle      string
		aggregateID string
		payload     json.RawMessage
		createdAt   time.Time
	}
)

func New(handle string, aggregateID string, payload json.RawMessage) *Event {
	return &Event{
		handle:      handle,
		aggregateID: aggregateID,
		payload:     payload,
		createdAt:   time.Now(),
	}
}

func (e *Event) Handle() string           { return e.handle }
func (e *Event) AggregateID() string      { return e.aggregateID }
func (e *Event) Payload() json.RawMessage { return e.payload }
func (e *Event) CreatedAt() time.Time     { return e.createdAt }
