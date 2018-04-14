package commandhandler

import "github.com/FelipeUmpierre/measures/pkg/business/command"

type (
	Handler interface {
		Handle(command.Command)
	}
)
