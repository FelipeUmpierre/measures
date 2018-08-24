package handler

import (
	"net/http"

	"github.com/FelipeUmpierre/measures/pkg/aggregator"
	"github.com/FelipeUmpierre/measures/pkg/domain"
	"github.com/FelipeUmpierre/measures/pkg/event"
	"github.com/FelipeUmpierre/measures/pkg/repository"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/gofrs/uuid"
	"github.com/json-iterator/go"
	"github.com/json-iterator/go/extra"
	log "github.com/rs/zerolog/log"
)

type (
	measureRepository interface {
		EventsFrom(string) ([]repository.Measures, error)
		Save(repository.MeasureEvent) error
	}
)

// SaveMeasures handler responsible to get the request payload and save it
func SaveMeasures(measureRepo measureRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		extra.SupportPrivateFields()
		measurePayload := event.MeasuresPayload{}
		if err := jsoniter.NewDecoder(r.Body).Decode(&measurePayload); err != nil {
			log.Error().Err(err).Msg(`failed to decode payload request`)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, `failed to decode payload request`)
			return
		}

		aggID := uuid.FromStringOrNil(
			r.Context().Value(`aggregate_id`).(string),
		)

		measureEvent := event.NewMeasureAddedEvent(aggID, measurePayload)
		if err := measureRepo.Save(measureEvent); err != nil {
			log.Error().Err(err).Msg(`failed to save the event`)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, `failed to save the event`)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

// GetMeasureState returns the current state of the measure for a user
func GetMeasureState(measureRepo measureRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		aggID := chi.URLParam(r, `aggregateID`)
		if aggID == `` {
			log.Error().Msg(`you must inform the user aggregate id`)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, `you must inform the user aggregate id`)
			return
		}

		events, err := measureRepo.EventsFrom(aggID)
		if err != nil {
			log.Error().Err(err).Msg(`failed to collect the event`)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, `failed to collect the events`)
			return
		}

		measure := domain.Measure{}
		for _, event := range events {
			aggregator.Apply(&measure, event)
		}

		extra.SupportPrivateFields()
		body, err := jsoniter.Marshal(measure)
		if err != nil {
			log.Error().Err(err).Msg(`failed to marshal the measure`)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, `failed to marshal the measure`)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func GetMeasures(measureRepo measureRepository) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		aggID := chi.URLParam(r, `aggregateID`)
		if aggID == `` {
			log.Error().Msg(`you must inform the user aggregate id`)
			w.WriteHeader(http.StatusBadRequest)
			render.JSON(w, r, `you must inform the user aggregate id`)
			return
		}

		events, err := measureRepo.EventsFrom(aggID)
		if err != nil {
			log.Error().Err(err).Msg(`failed to collect the event`)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, `failed to collect the events`)
			return
		}

		measures := []domain.Measure{}
		for _, event := range events {
			measure := domain.Measure{}
			aggregator.Apply(&measure, event)

			measures = append(measures, measure)
		}

		extra.SupportPrivateFields()
		body, err := jsoniter.Marshal(measures)
		if err != nil {
			log.Error().Err(err).Msg(`failed to marshal the measure`)
			w.WriteHeader(http.StatusInternalServerError)
			render.JSON(w, r, `failed to marshal the measure`)
			return
		}

		w.Write(body)
	}
}
