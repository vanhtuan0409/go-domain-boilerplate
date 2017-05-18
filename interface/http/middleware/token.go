package middleware

import (
	"context"
	"net/http"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

// ITokenProvider Sample token provider interface
// Should modify for correct type return
type ITokenParser interface {
	ParseToken(token string) (member.MemberID, error)
	ParseTokenFromHeader(r *http.Request) (member.MemberID, error)
}

type TokenMiddleware struct {
	parser ITokenParser
}

func NewTokenMiddleware(p ITokenParser) *TokenMiddleware {
	return &TokenMiddleware{
		parser: p,
	}
}

func (t *TokenMiddleware) ServeHTTP(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	memberID, err := t.parser.ParseTokenFromHeader(r)
	if err != nil {
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte(err.Error()))
		return
	}
	newContext := context.WithValue(r.Context(), "memberID", memberID)
	next.ServeHTTP(rw, r.WithContext(newContext))
}
