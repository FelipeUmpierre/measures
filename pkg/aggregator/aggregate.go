package aggregator

import (
	"github.com/FelipeUmpierre/measures/pkg/domain"
	"github.com/FelipeUmpierre/measures/pkg/event"
	"github.com/FelipeUmpierre/measures/pkg/repository"
)

func Apply(m *domain.Measure, e repository.Measures) {
	switch e.Handle {
	case event.MeasuresAdded:
		ev := event.NewMeasureEventFromStore(e)
		m.Added(ev)
	}
}
