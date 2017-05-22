package handler

import (
	"encoding/json"
	"net/http"

	"github.com/vanhtuan0409/go-domain-boilerplate/application/goal"
	goaldomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/goal"
	memberdomain "github.com/vanhtuan0409/go-domain-boilerplate/domain/member"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
)

type IErrorHandler interface {
	MapHttpCode(err error) int
	MapContent(err error) interface{}
}

type Response struct {
	HttpCode int
	Content  interface{}
}

func (r *Response) SendJSON(w http.ResponseWriter) error {
	js, err := json.Marshal(r.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(r.HttpCode)
	w.Write(js)
	return nil
}

type Builder struct {
	httpCode int
	content  interface{}
}

func ReponseBuilder() *Builder {
	builder := Builder{}
	return &builder
}

func (b *Builder) HttpCode(code int) *Builder {
	b.httpCode = code
	return b
}

func (b *Builder) Content(content interface{}) *Builder {
	b.content = content
	return b
}

func (b *Builder) Build() *Response {
	response := Response{}
	response.HttpCode = http.StatusNoContent
	if b.httpCode != 0 {
		response.HttpCode = b.httpCode
	}
	response.Content = ""
	if b.content != nil {
		response.Content = b.content
	}
	return &response
}

var ResponseOK = func(content interface{}) *Response {
	response := Response{}
	response.HttpCode = http.StatusOK
	response.Content = content
	return &response
}

var ResponseError = func(err error) *Response {
	response := Response{}
	response.Content = err.Error()
	if err == goaldomain.ErrorGoalNotFound {
		response.HttpCode = http.StatusNotFound
	} else if err == goaldomain.ErrorTaskNotFound {
		response.HttpCode = http.StatusNotFound
	} else if err == memberdomain.ErrorMemberNotFound {
		response.HttpCode = http.StatusNotFound
	} else if err == goal.ErrorUnauthorizeAccessGoal {
		response.HttpCode = http.StatusUnauthorized
	} else if err == ErrorInvalidToken {
		response.HttpCode = http.StatusUnauthorized
	} else if err == ErrorParseCheckInRequest {
		response.HttpCode = http.StatusBadRequest
	} else {
		response.HttpCode = http.StatusInternalServerError
	}
	logger.Logger.Error(response.Content)
	return &response
}

var ResponseFactory = func(content interface{}, err error) *Response {
	if err != nil {
		return ResponseError(err)
	}
	return ResponseOK(content)
}
