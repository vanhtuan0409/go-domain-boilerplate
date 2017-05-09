package goal

var (
	EventAddTaskToGoalType = "EventAddTaskToGoalType"
	EventCheckInTaskType   = "EventCheckInTaskType"
)

type EventAddTaskToGoal struct {
	GoalID    GoalID
	AddedTask *Task
}

func NewEventAddTaskToGoal(goalID GoalID, task *Task) *EventAddTaskToGoal {
	event := EventAddTaskToGoal{}
	event.GoalID = goalID
	event.AddedTask = task
	return &event
}

func (*EventAddTaskToGoal) GetEventType() string {
	return EventAddTaskToGoalType
}

type EventCheckInTask struct {
	GoalID   GoalID
	TaskName string
	OldValue int
	NewValue int
	Message  string
}

func NewEventCheckInTask(
	goalID GoalID, taskName string,
	oldValue, newValue int, message string,
) *EventCheckInTask {
	event := EventCheckInTask{}
	event.GoalID = goalID
	event.TaskName = taskName
	event.OldValue = oldValue
	event.NewValue = newValue
	event.Message = message
	return &event
}

func (*EventCheckInTask) GetEventType() string {
	return EventCheckInTaskType
}
