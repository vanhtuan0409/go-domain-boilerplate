package handler

import (
	"net/http"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	goaldomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	memberdomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
)

type ErrorMapper struct{}

func (*ErrorMapper) MapHttpCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if err == goaldomain.ErrorGoalNotFound {
		return http.StatusNotFound
	}
	if err == memberdomain.ErrorMemberNotFound {
		return http.StatusNotFound
	}
	if err == goal.ErrorUnauthorizeAccessGoal {
		return http.StatusUnauthorized
	}
	return http.StatusInternalServerError
}

func (*ErrorMapper) MapContent(err error) interface{} {
	return err.Error()
}
