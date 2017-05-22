package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	domaingoal "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	domainmember "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
	"github.com/vanhtuan0409/go-domain-boilerplate/interface/http/requestmodel"
)

var (
	ErrorParseCheckInRequest = errors.New("Parse Check in request failed")
	ErrorInvalidToken        = errors.New("Invalid Token")
)

type Controller struct {
	GoalUsecase   IGoalUsecase
	MemberUsecase IMemberUsecase
}

func NewController(gu IGoalUsecase, mu IMemberUsecase) *Controller {
	controller := Controller{}
	controller.GoalUsecase = gu
	controller.MemberUsecase = mu
	return &controller
}

func (ctrl *Controller) ListAllMember(w http.ResponseWriter, r *http.Request) {
	members, err := ctrl.MemberUsecase.ListAllMember()
	ResponseFactory(members, err).SendJSON(w)
}

func (ctrl *Controller) ListMemberGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := domainmember.MemberID(vars["memberID"])
	goals, err := ctrl.GoalUsecase.GetAllByMember(memberID)
	ResponseFactory(goals, err).SendJSON(w)
}

func (ctrl *Controller) CheckInTask(w http.ResponseWriter, r *http.Request) {
	// Parse goalID
	vars := mux.Vars(r)
	goalID := domaingoal.GoalID(vars["goalID"])

	// Parse memberID
	memberID := r.Context().Value("memberID").(domainmember.MemberID)

	// Parse checkin request
	decoder := json.NewDecoder(r.Body)
	checkin := requestmodel.CheckIn{}
	if err := decoder.Decode(&checkin); err != nil {
		ResponseError(ErrorParseCheckInRequest).SendJSON(w)
		return
	}

	// Excute checkin logic
	goal, err := ctrl.GoalUsecase.CheckInGoal(
		memberID, goalID,
		checkin.Name, checkin.Value, checkin.Message,
	)
	ResponseFactory(goal, err).SendJSON(w)
}
