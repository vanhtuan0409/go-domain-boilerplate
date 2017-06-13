package goal

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IGoalRepository interface {
	Get(goalID goal.GoalID) (*goal.Goal, error)
	Save(goal *goal.Goal) error
}

type IMemberRepository interface {
	Get(memberID member.MemberID) (*member.Member, error)
}

type IAccessControl interface {
	CanAccessGoal(member *member.Member, goal *goal.Goal) bool
}
