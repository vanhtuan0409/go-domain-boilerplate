package repository

import (
	"sync"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

var (
	defaultGoalDB = map[goal.GoalID]*goal.Goal{
		goal.GoalID("1"): &goal.Goal{
			ID:          goal.GoalID("1"),
			OwnerID:     member.MemberID("1"),
			Name:        "Goal 1",
			Description: "Meow meow",
			Tasks: []*goal.Task{
				&goal.Task{
					Name:          "Goal 1 Task 1",
					Description:   "",
					CurrentValue:  5,
					ExpectedValue: 10,
					Unit:          "customers",
				},
			},
		},
		goal.GoalID("2"): &goal.Goal{
			ID:          goal.GoalID("2"),
			OwnerID:     member.MemberID("1"),
			Name:        "Goal 2",
			Description: "Grao Grao",
			Tasks:       []*goal.Task{},
		},
		goal.GoalID("3"): &goal.Goal{
			ID:          goal.GoalID("3"),
			OwnerID:     member.MemberID("2"),
			Name:        "Goal 3",
			Description: "Cap cap",
			Tasks:       []*goal.Task{},
		},
		goal.GoalID("4"): &goal.Goal{
			ID:          goal.GoalID("4"),
			OwnerID:     member.MemberID("2"),
			Name:        "Goal 4",
			Description: "Blah blah",
			Tasks: []*goal.Task{
				&goal.Task{
					Name:          "Goal 4 Task 1",
					Description:   "Asasdasdasd",
					CurrentValue:  0,
					ExpectedValue: 100,
					Unit:          "%",
				},
			},
		},
	}
)

type InMemGoalRepo struct {
	mtx  sync.RWMutex
	goal map[goal.GoalID]*goal.Goal
}

func NewInMemGoalRepo() *InMemGoalRepo {
	return &InMemGoalRepo{
		goal: defaultGoalDB,
	}
}

func (r *InMemGoalRepo) Get(goalID goal.GoalID) (*goal.Goal, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if val, ok := r.goal[goalID]; ok {
		return val, nil
	}
	return nil, goal.ErrorGoalNotFound
}

func (r *InMemGoalRepo) GetAllByMember(memberID member.MemberID) ([]*goal.Goal, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	res := []*goal.Goal{}
	for _, val := range r.goal {
		if val.OwnerID == memberID {
			res = append(res, val)
		}
	}
	return res, nil
}
func (r *InMemGoalRepo) Save(goal *goal.Goal) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.goal[goal.ID] = goal
	return nil
}
