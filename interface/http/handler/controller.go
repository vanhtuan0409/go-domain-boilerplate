package handler

import (
	"net/http"

	domaingoal "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	domainmember "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type Controller struct {
	GoalUsecase   IGoalUsecase
	MemberUsecase IMemberUsecase
	Mapper        IErrorMapper
}

func NewController(gu IGoalUsecase, mu IMemberUsecase, mapper IErrorMapper) *Controller {
	controller := Controller{}
	controller.GoalUsecase = gu
	controller.MemberUsecase = mu
	controller.Mapper = mapper
	return &controller
}

func (ctrl *Controller) ListAllMember(w http.ResponseWriter, req *http.Request) {
	members, err := ctrl.MemberUsecase.ListAllMember()

	response := ReponseBuilder(ctrl.Mapper).Content(members).Error(err).Build()
	SendJSONResponse(response, w)
}

func (ctrl *Controller) ListMemberGoal(w http.ResponseWriter, req *http.Request) {
	memberID := domainmember.MemberID("1")
	goals, err := ctrl.GoalUsecase.GetAllByMember(memberID)

	response := ReponseBuilder(ctrl.Mapper).Content(goals).Error(err).Build()
	SendJSONResponse(response, w)
}

func (ctrl *Controller) CheckInTask(w http.ResponseWriter, req *http.Request) {
	memberID := domainmember.MemberID("1")
	goalID := domaingoal.GoalID("1")
	goal, err := ctrl.GoalUsecase.CheckInGoal(memberID, goalID, "Task 1", 50, "First check in")

	response := ReponseBuilder(ctrl.Mapper).Content(goal).Error(err).Build()
	SendJSONResponse(response, w)
}
