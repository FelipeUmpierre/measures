package handler

import (
	"net/http"

	"github.com/FelipeUmpierre/measures/pkg/business/command"
	"github.com/go-chi/render"
	log "github.com/sirupsen/logrus"
)

type (
	cmd interface {
		CommandType() string
	}

	dispatcher interface {
		Dispatch(command.Command)
	}
)

const (
	MeasureCreated = `measure.created`
)

func SaveMeasure(d dispatcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		command := r.Context().Value(`measure`).(*command.MeasureCommand)
		userID := r.Context().Value(`userID`).(string)

		d.Dispatch(command)

		log.WithFields(log.Fields{
			`command`: command,
			`userID`:  userID,
		}).Infoln(`Measure saved with success!`)
		w.WriteHeader(http.StatusOK)
		render.JSON(w, r, ErrorResponse{Message: `Measure saved with success!`})
	}
}
