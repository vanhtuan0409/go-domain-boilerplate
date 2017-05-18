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
	Mapper        IErrorHandler
}

func NewController(gu IGoalUsecase, mu IMemberUsecase, mapper IErrorHandler) *Controller {
	controller := Controller{}
	controller.GoalUsecase = gu
	controller.MemberUsecase = mu
	controller.Mapper = mapper
	return &controller
}

func (ctrl *Controller) ListAllMember(w http.ResponseWriter, r *http.Request) {
	members, err := ctrl.MemberUsecase.ListAllMember()

	response := ReponseBuilder(ctrl.Mapper).Content(members).Error(err).Build()
	SendJSONResponse(response, w)
}

func (ctrl *Controller) ListMemberGoal(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	memberID := domainmember.MemberID(vars["memberID"])
	goals, err := ctrl.GoalUsecase.GetAllByMember(memberID)

	response := ReponseBuilder(ctrl.Mapper).Content(goals).Error(err).Build()
	SendJSONResponse(response, w)
}

func (ctrl *Controller) CheckInTask(w http.ResponseWriter, r *http.Request) {
	// Parse goalID
	vars := mux.Vars(r)
	goalID := domaingoal.GoalID(vars["goalID"])

	// Parse memberID
	token := r.Header.Get("Authorization")
	if len(token) <= 7 {
		response := ReponseBuilder(ctrl.Mapper).Error(ErrorInvalidToken).Build()
		SendJSONResponse(response, w)
		return
	}
	memberID := domainmember.MemberID(token[7:])

	// Parse checkin request
	decoder := json.NewDecoder(r.Body)
	checkin := requestmodel.CheckIn{}
	if err := decoder.Decode(&checkin); err != nil {
		response := ReponseBuilder(ctrl.Mapper).Error(ErrorParseCheckInRequest).Build()
		SendJSONResponse(response, w)
		return
	}

	// Excute checkin logic
	goal, err := ctrl.GoalUsecase.CheckInGoal(
		memberID, goalID,
		checkin.Name, checkin.Value, checkin.Message,
	)

	// Return response
	response := ReponseBuilder(ctrl.Mapper).Content(goal).Error(err).Build()
	SendJSONResponse(response, w)
}
