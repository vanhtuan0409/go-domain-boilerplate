package http

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	applicationgoal "github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

var (
	AUTH_CONTEXT_KEY = "authInfo"
)

type ITokenParser interface {
	ParseToken(token string) (*accesscontrol.AuthInfo, error)
	ParseTokenFromHeader(r *http.Request) (*accesscontrol.AuthInfo, error)
}

func TokenMiddleware(next http.Handler, parser ITokenParser) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authInfo, err := parser.ParseTokenFromHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Authorization header format must be Bearer {token}"))
			server.Log.Error("Authorization header format must be Bearer {token}")
			return
		}
		newContext := context.WithValue(r.Context(), AUTH_CONTEXT_KEY, authInfo)
		next.ServeHTTP(w, r.WithContext(newContext))
	})
}

func ErrorHandle(next server.JSONContextEndpoint) server.JSONContextEndpoint {
	return func(c context.Context, r *http.Request) (int, interface{}, error) {
		code, res, err := next(c, r)
		if err != nil {
			server.Log.Error("Unexpected error: ", err.Error())
			if err == goal.ErrorGoalNotFound {
				return http.StatusNotFound, "", nil
			} else if err == goal.ErrorTaskNotFound {
				return http.StatusNotFound, "", nil
			} else if err == member.ErrorMemberNotFound {
				return http.StatusNotFound, "", nil
			} else if err == applicationgoal.ErrorUnauthorizeAccessGoal {
				return http.StatusUnauthorized, "", nil
			} else if err == ErrorInvalidToken {
				return http.StatusUnauthorized, "", nil
			} else if err == ErrorParseCheckInRequest {
				return http.StatusBadRequest, "", nil
			} else {
				return http.StatusInternalServerError, "", nil
			}
		}

		return code, res, err
	}
}
