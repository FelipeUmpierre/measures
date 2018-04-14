package bus

import (
	"sync"

	"github.com/FelipeUmpierre/measures/pkg/business/command"
	"github.com/FelipeUmpierre/measures/pkg/business/handler"
)

type (
	// Bus ...
	Bus struct {
		mu       *sync.RWMutex
		handlers map[string]commandhandler.Handler
	}
)

// New ...
func New() *Bus {
	return &Bus{}
}

// Register ...
func (b *Bus) Register(cmd command.Command, h commandhandler.Handler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[cmd.CommandType()] = h
}

// Dispatch ...
func (b *Bus) Dispatch(cmd command.Command) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	h, ok := b.handlers[cmd.CommandType()]
	if !ok {
		return
	}

	h.Handle(cmd)
}
