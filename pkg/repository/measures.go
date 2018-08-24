package repository

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	lk "github.com/ulule/loukoum"
)

type (
	// MeasuresRepo holds the database connection
	MeasuresRepo struct {
		db *sqlx.DB
	}

	Measures struct {
		UUID        string `db:"uuid"`
		Handle      string `db:"handle"`
		AggregateID string `db:"aggregate_id"`
		Payload     []byte `db:"payload"`
	}

	MeasureEvent interface {
		UUID() string
		Handle() string
		AggregateID() string
		Payload() json.Marshaler
	}
)

// NewMeasuresRepository creates a new repository
func NewMeasuresRepository(db *sqlx.DB) *MeasuresRepo {
	return &MeasuresRepo{db}
}

// EventsFrom search for all events from a specific aggregate_id
func (m *MeasuresRepo) EventsFrom(ID string) ([]Measures, error) {
	query, args := lk.Select(`uuid`, `handle`, `aggregate_id`, `payload`).
		From(`events`).
		Where(lk.Condition(`aggregate_id`).Equal(ID)).
		NamedQuery()

	stmt, err := m.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	measures := []Measures{}
	err = stmt.Select(&measures, args)

	return measures, err
}

// Save ...
func (m *MeasuresRepo) Save(e MeasureEvent) error {
	payload, err := json.Marshal(e.Payload())
	if err != nil {
		return err
	}

	query, args := lk.Insert(`events`).
		Set(
			lk.Pair(`uuid`, e.UUID()),
			lk.Pair(`handle`, e.Handle()),
			lk.Pair(`aggregate_id`, e.AggregateID()),
			lk.Pair(`payload`, payload),
			lk.Pair(`created_at`, time.Now().UTC().Format(time.RFC3339)),
		).
		NamedQuery()

	stmt, err := m.db.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args)
	return err
}
