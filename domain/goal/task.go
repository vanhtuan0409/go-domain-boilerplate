package goal

type Task struct {
	Name          string
	Description   string
	CurrentValue  int
	ExpectedValue int
	Unit          string
}

func NewTask(
	name string, description string,
	expectedValue int, unit string,
) *Task {
	task := Task{}
	task.Name = name
	task.Description = description
	task.CurrentValue = 0
	task.ExpectedValue = expectedValue
	task.Unit = unit
	return &task
}

func (t *Task) Progress() int {
	progress := (t.CurrentValue * 100) / t.ExpectedValue
	if progress > 100 {
		return 100
	}
	return progress
}
