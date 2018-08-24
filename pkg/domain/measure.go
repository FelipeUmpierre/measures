package domain

import (
	"github.com/FelipeUmpierre/measures/pkg/event"
)

type Measure struct {
	chest float64
}

func (m *Measure) Added(e event.Measures) {
	payload := e.MeasuresPayload

	m.chest = payload.Chest()
}
