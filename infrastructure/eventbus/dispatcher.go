package eventbus

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/common"
)

type Dispatcher struct{}

func NewEventDispatcher() *Dispatcher {
	return &Dispatcher{}
}

func (d *Dispatcher) Dispatch(events ...common.IDomainEvent) {
	// TODO
}
