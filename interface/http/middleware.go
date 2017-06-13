package http

import (
	"context"
	"net/http"

	"github.com/NYTimes/gizmo/server"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
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
