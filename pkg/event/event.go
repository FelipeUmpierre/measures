package event

import uuid "github.com/gofrs/uuid"

// Event holds the default schema for the payload
type Event struct {
	uuid        uuid.UUID
	handle      string
	aggregateID uuid.UUID
}

// NewEvent creates the event skeleton
func NewEvent(aggregateID uuid.UUID, handle string) Event {
	id, _ := uuid.NewV4()

	return Event{
		uuid:        id,
		handle:      handle,
		aggregateID: aggregateID,
	}
}

// UUID returns the uuid from the event
func (e Event) UUID() string { return e.uuid.String() }

// Handle returns the handle from the event
func (e Event) Handle() string { return e.handle }

// AggregateID returns the aggregate id from the event
func (e Event) AggregateID() string { return e.aggregateID.String() }
