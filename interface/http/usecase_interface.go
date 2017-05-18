package http

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IGoalUsecase interface {
	GetAllByMember(memberID member.MemberID) ([]*goal.Goal, error)
	CheckInGoal(
		actorID member.MemberID, goalID goal.GoalID,
		taskName string, newValue int, message string,
	) (*goal.Goal, error)
}

type IMemberUsecase interface {
	ListAllMember() ([]*member.Member, error)
}
