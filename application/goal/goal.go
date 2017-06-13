package goal

import (
	"errors"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/common"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

var (
	ErrorUnauthorizeAccessGoal = errors.New("Unauthorize Access to Goal")
)

type GoalUsecase struct {
	GoalRepo      IGoalRepository
	MemberRepo    IMemberRepository
	AccessControl IAccessControl
	Dispatcher    common.IEventDispatcher
}

func NewGoalUsecase(
	goalRepo IGoalRepository, memberRepo IMemberRepository,
	acessControl IAccessControl, dispatcher common.IEventDispatcher,
) *GoalUsecase {
	usecase := GoalUsecase{}
	usecase.GoalRepo = goalRepo
	usecase.MemberRepo = memberRepo
	usecase.AccessControl = acessControl
	usecase.Dispatcher = dispatcher
	return &usecase
}

func (u *GoalUsecase) AddTaskToGoal(
	actorID member.MemberID, goalID goal.GoalID,
	taskName, description string,
	expectedValue int, unit string,
) (*goal.Goal, error) {
	actor, err := u.MemberRepo.Get(actorID)
	if err != nil {
		return nil, err
	}
	goal, err := u.GoalRepo.Get(goalID)
	if err != nil {
		return nil, err
	}
	if !u.AccessControl.CanAccessGoal(actor, goal) {
		return nil, ErrorUnauthorizeAccessGoal
	}

	if err = goal.AddTask(taskName, description, expectedValue, unit); err != nil {
		return nil, err
	}
	if err = u.GoalRepo.Save(goal); err != nil {
		return nil, err
	}

	u.Dispatcher.Dispatch(goal.Events...)
	return goal, nil
}

func (u *GoalUsecase) CheckInGoal(
	actorID member.MemberID, goalID goal.GoalID,
	taskName string, newValue int, message string,
) (*goal.Goal, error) {
	actor, err := u.MemberRepo.Get(actorID)
	if err != nil {
		return nil, err
	}
	goal, err := u.GoalRepo.Get(goalID)
	if err != nil {
		return nil, err
	}
	if !u.AccessControl.CanAccessGoal(actor, goal) {
		return nil, ErrorUnauthorizeAccessGoal
	}

	if err = goal.CheckIn(taskName, newValue, message); err != nil {
		return nil, err
	}
	if err = u.GoalRepo.Save(goal); err != nil {
		return nil, err
	}

	// u.Dispatcher.Dispatch(goal.Events...)
	return goal, nil
}
