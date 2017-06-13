package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/gorilla/mux"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	domaingoal "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	domainmember "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IGoalUsecase interface {
	CheckInGoal(
		actorID domainmember.MemberID, goalID domaingoal.GoalID,
		taskName string, newValue int, message string,
	) (*domaingoal.Goal, error)
}

type GoalEndPoints struct {
	gu          IGoalUsecase
	tokenparser ITokenParser
}

func NewGoalEndPoints(gu IGoalUsecase, p ITokenParser) *GoalEndPoints {
	endpoints := GoalEndPoints{}
	endpoints.gu = gu
	endpoints.tokenparser = p
	return &endpoints
}

func (e *GoalEndPoints) Prefix() string {
	return "/api/goals"
}

func (e *GoalEndPoints) Middleware(h http.Handler) http.Handler {
	chained := server.CORSHandler(h, "*")
	chained = TokenMiddleware(chained, e.tokenparser)
	return chained
}

func (e *GoalEndPoints) ContextMiddleware(h server.ContextHandler) server.ContextHandler {
	return h
}

func (e *GoalEndPoints) JSONContextMiddleware(h server.JSONContextEndpoint) server.JSONContextEndpoint {
	return ErrorHandle(h)
}

func (e *GoalEndPoints) ContextEndpoints() map[string]map[string]server.ContextHandlerFunc {
	return map[string]map[string]server.ContextHandlerFunc{}
}

func (e *GoalEndPoints) JSONEndpoints() map[string]map[string]server.JSONContextEndpoint {
	return map[string]map[string]server.JSONContextEndpoint{
		"/{goalID}/checkin": map[string]server.JSONContextEndpoint{
			"POST": e.CheckIn,
		},
	}
}

type CheckIn struct {
	Name    string `json:"name"`
	Value   int    `json:"value"`
	Message string `json:"message"`
}

func (e *GoalEndPoints) CheckIn(c context.Context, r *http.Request) (int, interface{}, error) {
	// Parse goalID
	vars := mux.Vars(r)
	goalID := domaingoal.GoalID(vars["goalID"])

	// Parse checkin request
	decoder := json.NewDecoder(r.Body)
	checkin := CheckIn{}
	if err := decoder.Decode(&checkin); err != nil {
		return 0, "", ErrorParseCheckInRequest
	}

	// Parse auth from context
	raw := r.Context().Value("authInfo")
	authInfo, ok := raw.(*accesscontrol.AuthInfo)
	if !ok {
		return 0, "", ErrorInvalidToken
	}

	// Excute checkin logic
	goal, err := e.gu.CheckInGoal(
		authInfo.MemberID, goalID,
		checkin.Name, checkin.Value, checkin.Message,
	)
	return 0, goal, err
}
