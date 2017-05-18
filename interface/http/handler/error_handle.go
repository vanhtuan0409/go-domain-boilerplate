package handler

import (
	"net/http"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	goaldomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	memberdomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
)

type ErrorMapper struct{}

func (*ErrorMapper) MapHttpCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if err == goaldomain.ErrorGoalNotFound {
		return http.StatusNotFound
	}
	if err == goaldomain.ErrorTaskNotFound {
		return http.StatusNotFound
	}
	if err == memberdomain.ErrorMemberNotFound {
		return http.StatusNotFound
	}
	if err == goal.ErrorUnauthorizeAccessGoal {
		return http.StatusUnauthorized
	}
	if err == ErrorInvalidToken {
		return http.StatusUnauthorized
	}
	if err == ErrorParseCheckInRequest {
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func (*ErrorMapper) MapContent(err error) interface{} {
	logger.Logger.Error(err.Error())
	return err.Error()
}
