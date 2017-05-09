package common

type BaseEntity struct {
	Events []IDomainEvent
}

type IDomainEvent interface {
	GetEventType() string
}

func (e *BaseEntity) IsContainEvent(eventType string) bool {
	for _, evt := range e.Events {
		if evt.GetEventType() == eventType {
			return true
		}
	}
	return false
}

func (e *BaseEntity) DeferEvent(evt IDomainEvent) {
	evtType := evt.GetEventType()
	if !e.IsContainEvent(evtType) {
		e.Events = append(e.Events)
	}
}
