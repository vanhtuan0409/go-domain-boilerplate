package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/gorilla/mux"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IGoalUsecase interface {
	CheckInGoal(
		actorID member.MemberID, goalID goal.GoalID,
		taskName string, newValue int, message string,
	) (*goal.Goal, error)
}

type GoalEndPoints struct {
	gu IGoalUsecase
}

func NewGoalEndPoints(gu IGoalUsecase) *GoalEndPoints {
	endpoints := GoalEndPoints{}
	endpoints.gu = gu
	return &endpoints
}

func (e *GoalEndPoints) Prefix() string {
	return "/api/goals"
}

func (e *GoalEndPoints) Middleware(h http.Handler) http.Handler {
	return h
}

func (e *GoalEndPoints) ContextMiddleware(h server.ContextHandler) server.ContextHandler {
	return h
}

func (e *GoalEndPoints) JSONContextMiddleware(h server.JSONContextEndpoint) server.JSONContextEndpoint {
	return h
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
	goalID := goal.GoalID(vars["goalID"])

	// Parse checkin request
	decoder := json.NewDecoder(r.Body)
	checkin := CheckIn{}
	if err := decoder.Decode(&checkin); err != nil {
		return 0, "", ErrorParseCheckInRequest
	}

	// Parse auth from context
	memberID := member.MemberID("1")

	// Excute checkin logic
	goal, err := e.gu.CheckInGoal(
		memberID, goalID,
		checkin.Name, checkin.Value, checkin.Message,
	)
	return 0, goal, err
}
