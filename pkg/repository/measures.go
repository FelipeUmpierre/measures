package repository

import (
	"github.com/FelipeUmpierre/measures/pkg/domain"
	"github.com/jmoiron/sqlx"
	lk "github.com/ulule/loukoum"
)

type (
	// MeasuresRepo holds the database connection
	MeasuresRepo struct {
		db *sqlx.DB
	}
)

// NewMeasuresRepository creates a new repository
func NewMeasuresRepository(db *sqlx.DB) *MeasuresRepo {
	return &MeasuresRepo{db}
}

// EventsFrom search for all events from a specific aggregate_id
func (m *MeasuresRepo) EventsFrom(ID string) (*[]domain.Measures, error) {
	query, args := lk.Select(`id`, `handle`, `aggregate_id`, `payload`).
		From(`measures`).
		Where(lk.Condition(`aggregate_id`).Equal(ID)).
		Prepare()

	stmt, err := m.db.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	measures := new([]domain.Measures)
	err = stmt.Select(measures, args)

	return measures, err
}
