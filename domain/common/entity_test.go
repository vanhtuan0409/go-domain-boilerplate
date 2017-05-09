package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockEvent struct {
	evtType string
}

func (e *mockEvent) GetEventType() string {
	return e.evtType
}

func TesBaseEntity(t *testing.T) {
	entity := BaseEntity{}
	assert.Equal(t, 0, len(entity.Events))
	assert.False(t, entity.IsContainEvent("Event 1"))

	evt1 := &mockEvent{"Event 1"}
	entity.DeferEvent(evt1)
	assert.Equal(t, []IDomainEvent{evt1}, entity.Events)
	assert.True(t, entity.IsContainEvent("Event 1"))

	entity.DeferEvent(evt1)
	assert.Equal(t, []IDomainEvent{evt1}, entity.Events)
	assert.True(t, entity.IsContainEvent("Event 1"))

	evt2 := &mockEvent{"Event 2"}
	entity.DeferEvent(evt2)
	assert.Equal(t, []IDomainEvent{evt1, evt2}, entity.Events)
	assert.True(t, entity.IsContainEvent("Event 1"))
	assert.True(t, entity.IsContainEvent("Event 2"))
}
