package event

import (
	"encoding/json"

	"github.com/FelipeUmpierre/go/extra"
	"github.com/FelipeUmpierre/measures/pkg/repository"
	uuid "github.com/gofrs/uuid"
	jsoniter "github.com/json-iterator/go"
)

type (
	// Measures hold the event struct
	Measures struct {
		Event
		MeasuresPayload
	}

	// MeasuresPayload contains the payload schema
	MeasuresPayload struct {
		chest float64 `json:"chest"`
	}
)

// NewMeasureAddedEvent creates a new Measure event with event default data and payload
func NewMeasureAddedEvent(
	aggregateID uuid.UUID,
	payload MeasuresPayload,
) Measures {
	return Measures{
		Event:           NewEvent(aggregateID, MeasuresAdded),
		MeasuresPayload: payload,
	}
}

// NewMeasureEventFromStore creates a new event
func NewMeasureEventFromStore(m repository.Measures) Measures {
	extra.SupportPrivateFields()
	payload := MeasuresPayload{}
	jsoniter.Unmarshal(m.Payload, &payload)

	return Measures{
		Event: Event{
			aggregateID: uuid.FromStringOrNil(m.AggregateID),
			handle:      m.Handle,
			uuid:        uuid.FromStringOrNil(m.UUID),
		},
		MeasuresPayload: payload,
	}
}

// Payload returns the json.Marshaler interface
func (m Measures) Payload() json.Marshaler { return m.MeasuresPayload }

// MarshalJSON custom json marshaler
func (m MeasuresPayload) MarshalJSON() ([]byte, error) {
	type payload MeasuresPayload
	return jsoniter.Marshal((payload)(m))
}

// Chest returns the chest value from the request payload
func (m MeasuresPayload) Chest() float64 { return m.chest }
