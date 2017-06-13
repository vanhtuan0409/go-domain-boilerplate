package http

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type IMemberUsecase interface {
	ListAllMember() ([]*member.Member, error)
	// GetMemberGoals(memberID member.MemberID) ([]*goal.Goal, error)
}

type MemberEndPoints struct {
	mu IMemberUsecase
}

func NewMemberEndPoints(mu IMemberUsecase) *MemberEndPoints {
	endpoints := MemberEndPoints{}
	endpoints.mu = mu
	return &endpoints
}

func (e *MemberEndPoints) Prefix() string {
	return "/api/members"
}

func (e *MemberEndPoints) Middleware(h http.Handler) http.Handler {
	return h
}

func (e *MemberEndPoints) ContextMiddleware(h server.ContextHandler) server.ContextHandler {
	return h
}

func (e *MemberEndPoints) JSONContextMiddleware(h server.JSONContextEndpoint) server.JSONContextEndpoint {
	return h
}

func (e *MemberEndPoints) ContextEndpoints() map[string]map[string]server.ContextHandlerFunc {
	return map[string]map[string]server.ContextHandlerFunc{}
}

func (e *MemberEndPoints) JSONEndpoints() map[string]map[string]server.JSONContextEndpoint {
	return map[string]map[string]server.JSONContextEndpoint{
		"": map[string]server.JSONContextEndpoint{
			"GET": e.Members,
		},
		// "/{memberID}/goals": map[string]server.JSONContextEndpoint{
		// 	"GET": e.MemberGoals,
		// },
	}
}

func (e *MemberEndPoints) Members(c context.Context, r *http.Request) (int, interface{}, error) {
	members, err := e.mu.ListAllMember()
	return 0, members, err
}

// func (e *MemberEndPoints) MemberGoals(c context.Context, r *http.Request) (int, interface{}, error) {
// 	vars := mux.Vars(r)
// 	memberID := member.MemberID(vars["memberID"])
// 	goals, err := e.mu.GetMemberGoals(memberID)
// 	return 0, goals, err
// }
