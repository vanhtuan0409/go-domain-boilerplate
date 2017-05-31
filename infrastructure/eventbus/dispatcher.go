package eventbus

import (
	"encoding/json"
	"fmt"

	"time"

	"github.com/nsqio/go-nsq"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/common"
	eventhandler "github.com/vanhtuan0409/go-domain-boilerplate/interface/eventbus"
)

type Dispatcher struct {
	nsqProducer *nsq.Producer
}

func NewEventDispatcher(p *nsq.Producer) *Dispatcher {
	dispatcher := Dispatcher{}
	dispatcher.nsqProducer = p
	return &dispatcher
}

func (d *Dispatcher) Dispatch(events ...common.IDomainEvent) error {
	for _, e := range events {
		err := d.dispatch(e)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Dispatcher) dispatch(event common.IDomainEvent) error {
	js, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return d.nsqProducer.Publish(event.GetEventType(), js)
}

type EventHandlerFunc func(eventhandler.IEventMessage) error
type eventMessage struct {
	e *nsq.Message
}

func (m *eventMessage) Body() []byte {
	return m.e.Body
}
func (m *eventMessage) Timestamp() time.Time {
	return time.Unix(0, m.e.Timestamp)
}
func (m *eventMessage) Attempts() uint16 {
	return m.e.Attempts
}
func (m *eventMessage) Retry() {
	if m.Attempts() < 5 {
		duration := time.Duration(5 * time.Minute)
		m.e.Requeue(duration)
	} else {
		fmt.Println("Retry message error")
	}
}

func MakeEventHandlerFunc(h EventHandlerFunc) nsq.HandlerFunc {
	return func(message *nsq.Message) error {
		m := eventMessage{message}
		return h(&m)
	}
}
