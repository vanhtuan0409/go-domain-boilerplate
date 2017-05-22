package tokenprovider

import (
	"errors"
	"net/http"
	"strings"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
	"github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type TokenProvider struct {
}

func NewTokenProvider() *TokenProvider {
	return &TokenProvider{}
}

func (t *TokenProvider) ParseToken(token string) (*accesscontrol.AuthInfo, error) {
	memberID := member.MemberID(token)
	auth := accesscontrol.AuthInfo{}
	auth.MemberID = memberID
	return &auth, nil
}

func (t *TokenProvider) ParseTokenFromHeader(r *http.Request) (*accesscontrol.AuthInfo, error) {
	token, err := t.getTokenFromHeader(r)
	if err != nil {
		return nil, err
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
