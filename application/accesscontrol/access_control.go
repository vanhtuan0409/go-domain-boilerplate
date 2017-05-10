package accesscontrol

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type AccessControl struct{}

func NewAccessControl() *AccessControl {
	accessControl := AccessControl{}
	return &accessControl
}

func (*AccessControl) CanAccessGoal(member *member.Member, goal *goal.Goal) bool {
	return member.ID == goal.OwnerID
}
