package common

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/common"
)

type IEventDispatcher interface {
	Dispatch(events ...common.IDomainEvent)
}
