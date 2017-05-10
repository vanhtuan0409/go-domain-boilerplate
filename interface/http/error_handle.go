package http

import (
	"net/http"
	"reflect"

    "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
    "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

func GetStatusCodeFromError(err error) int {
    if err == nil {
        return http.StatusOK
    }
    if err == goal.ErrorGoalNotFound {
        return http.StatusNotFound
    }
    if err == member.ErrorMemberNotFound {
        return http.StatusNotFound
    }
    return http.StatusInternalServerError
}
