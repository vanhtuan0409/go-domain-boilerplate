package http

import (
	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/member"
    domaingoal "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
    domainmember "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type Controller struct {
	GoalUsecase   goal.GoalUsecase
	MemberUsecase member.MemberUsecase
}

func (ctrl *Controller) RegisterMember(res http.ResponseWriter, req *http.Request) {
    // TODO: parse information from http request
    userName := "Shiro"
    user, err := ctrl.MemberUsecase.RegisterMember(userName)
}

func (ctrl *Controller) RegisterNewEmail(res http.ResponseWriter, req *http.Request) {
    // TODO: parse information from http request
    memberID := domainmember.MemberID("1")
    email := domainmember.Email("vanhtuan0409@gmail.com")
    member, err := ctrl.MemberUsecase.RegisterNewEmail(memberID, email)
}

func (ctrl *Controller) AddTaskToGoal(res http.ResponseWriter, req *http.Request) {
    // TODO: parse information from http request
    memberID := domainmember.MemberID("1")
    goalID := domaingoal.GoalID("1")
    goal, err := ctrl.GoalUsecase.AddTaskToGoal(memberID, goalID, "Task 1", "", 100, "time")
}

func (ctrl *Controller) CheckInTask(res http.ResponseWriter, req *http.Request) {
    // TODO: parse information from http request
    memberID := domainmember.MemberID("1")
    goalID := domaingoal.GoalID("1")
    goal, err := ctrl.GoalUsecase.CheckInGoal(memberID, goalID, "Task 1", 50, "First check in")
}
