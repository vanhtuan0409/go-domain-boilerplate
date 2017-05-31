package eventbus

import (
	"time"
)

type IEventMessage interface {
	Body() []byte
	Timestamp() time.Time
	Attempts() uint16
	Retry()
}
