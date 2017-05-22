package middleware

import (
	"context"
	"net/http"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
)

type ITokenParser interface {
	ParseToken(token string) (*accesscontrol.AuthInfo, error)
	ParseTokenFromHeader(r *http.Request) (*accesscontrol.AuthInfo, error)
}

type TokenMiddleware struct {
	key    string
	parser ITokenParser
}

func NewTokenMiddleware(k string, p ITokenParser) *TokenMiddleware {
	return &TokenMiddleware{
		parser: p,
		key:    k,
	}
}

func (t *TokenMiddleware) ServeHTTP(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	authInfo, err := t.parser.ParseTokenFromHeader(r)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("Authorization header format must be Bearer {token}"))
		logger.Logger.Error("Authorization header format must be Bearer {token}")
		return
	}
	newContext := context.WithValue(r.Context(), t.key, authInfo)
	next.ServeHTTP(rw, r.WithContext(newContext))
}
