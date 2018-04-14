package commandhandler

import (
	"github.com/FelipeUmpierre/measures/pkg/business/command"
	"github.com/FelipeUmpierre/measures/pkg/repository"
	log "github.com/sirupsen/logrus"
)

type (
	MeasureHandler struct {
		eventRepo eventRepo
	}

	eventRepo interface {
		Emit(repository.Event) error
	}
)

func NewMeasureHandler(repo eventRepo) *MeasureHandler {
	return &MeasureHandler{repo}
}

func (m *MeasureHandler) Handle(c command.Command) {
	log.Println(`oioio`)
}
