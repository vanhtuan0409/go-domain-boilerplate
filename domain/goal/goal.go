package goal

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/common"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type GoalID string

var (
	ErrorDuplicateTask    = errors.New("Duplicate task")
	ErrorInvalidTaskValue = errors.New("Invalid task value")
	ErrorTaskNotFound     = errors.New("Task not found")
)

type Goal struct {
	common.BaseEntity
	ID          GoalID
	OwnerID     member.MemberID
	Name        string
	Description string
	Tasks       []*Task
}

func NewGoal(name string, description string, owner *member.Member) *Goal {
	goal := Goal{}
	uid := uuid.NewV4().String()
	goal.ID = GoalID(uid)
	goal.Name = name
	goal.Description = description
	goal.OwnerID = owner.ID
	goal.Tasks = []*Task{}
	return &goal
}

func (g *Goal) IsHaveTask(taskName string) bool {
	for _, task := range g.Tasks {
		if task.Name == taskName {
			return true
		}
	}
	return false
}

func (g *Goal) AddTask(
	name string, description string,
	expectedValue int, unit string,
) error {
	if expectedValue <= 0 {
		return ErrorInvalidTaskValue
	}
	if g.IsHaveTask(name) {
		return ErrorDuplicateTask
	}
	task := NewTask(name, description, expectedValue, unit)
	g.Tasks = append(g.Tasks, task)

	event := NewEventAddTaskToGoal(g.ID, task)
	g.DeferEvent(event)

	return nil
}

func (g *Goal) FindTask(taskName string) (*Task, error) {
	for _, task := range g.Tasks {
		if task.Name == taskName {
			return task, nil
		}
	}
	return nil, ErrorTaskNotFound
}

func (g *Goal) CheckIn(taskName string, newValue int, message string) error {
	if newValue <= 0 {
		return ErrorInvalidTaskValue
	}

	task, err := g.FindTask(taskName)
	if err != nil {
		return err
	}

	oldTaskValue := task.CurrentValue
	task.CurrentValue = newValue

	event := NewEventCheckInTask(g.ID, taskName, oldTaskValue, task.CurrentValue, message)
	g.DeferEvent(event)

	return nil
}

func (g *Goal) Progress() int {
	if len(g.Tasks) == 0 {
		return 0
	}

	total := 0
	for _, task := range g.Tasks {
		total += task.Progress()
	}
	return total / len(g.Tasks)
}
