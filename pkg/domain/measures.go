package domain

import (
	"encoding/json"

	"github.com/satori/go.uuid"
)

type (
	Measures struct {
		ID          uuid.UUID       `db:"id" json:"id"`
		Handle      string          `db:"handle" json:"handle"`
		AggregateID uuid.UUID       `db:"aggregate_id" json:"aggregate_id"`
		Payload     json.RawMessage `db:"payload" json:"payload"`
	}
)
