package tokenprovider

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type TokenProvider struct {
}

func NewTokenProvider() *TokenProvider {
	return &TokenProvider{}
}

func (t *TokenProvider) ParseToken(token string) (member.MemberID, error) {
	return member.MemberID(token), nil
}

func (t *TokenProvider) ParseTokenFromHeader(r *http.Request) (member.MemberID, error) {
	token, err := t.getTokenFromHeader(r)
	if err != nil {
		return "", err
	}
	return t.ParseToken(token)
}

func (t *TokenProvider) getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
		return "", errors.New("Authorization header format must be Bearer {token}")
	}
	return authHeaderParts[1], nil
}
