package repository

import (
	"sync"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type InMemGoalRepo struct {
	mtx  sync.RWMutex
	goal map[goal.GoalID]*goal.Goal
}

func NewInMemGoalRepo() *InMemGoalRepo {
	return &InMemGoalRepo{
		goal: make(map[goal.GoalID]*goal.Goal),
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
