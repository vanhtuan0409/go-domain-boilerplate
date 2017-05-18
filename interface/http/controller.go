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

func (ctrl *Controller) ListAllMember(w http.ResponseWriter, req *http.Request) {
    members, err := ctrl.MemberUsecase.ListAllMember()

    response := ReponseBuilder().Content(members).Error(err).Build()
    SendJSONResponse(response, w)
}

func (ctrl *Controller) ListMemberGoal(w http.ResponseWriter, req *http.Request) {
    memberID := domainmember.MemberID("1")
    goals, err := ctrl.GoalUsecase.GetAllByMember(memberID)

    response := ReponseBuilder().Content(goals).Error(err).Build()
    SendJSONResponse(response, w)
}

func (ctrl *Controller) CheckInTask(w http.ResponseWriter, req *http.Request) {
    memberID := domainmember.MemberID("1")
    goalID := domaingoal.GoalID("1")
    goal, err := ctrl.GoalUsecase.CheckInGoal(memberID, goalID, "Task 1", 50, "First check in")

    response := ReponseBuilder().Content(goal).Error(err)
    SendJSONResponse(response, w)
}
