package handler

import (
	"encoding/json"
	"net/http"
)

type IErrorHandler interface {
	MapHttpCode(err error) int
	MapContent(err error) interface{}
}

type Response struct {
	HttpCode int
	Content  interface{}
	Err      error
}

type Builder struct {
	httpCode    int
	content     interface{}
	err         error
	errorMapper IErrorHandler
}

func ReponseBuilder(mapper IErrorHandler) *Builder {
	builder := Builder{}
	builder.errorMapper = mapper
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

func (b *Builder) Error(err error) *Builder {
	b.err = err
	return b
}

func (b *Builder) Build() *Response {
	response := Response{}
	response.HttpCode = b.getStatusCode()
	response.Content = b.getContent()
	response.Err = b.err
	return &response
}

func (b *Builder) getStatusCode() int {
	if b.err != nil {
		return b.errorMapper.MapHttpCode(b.err)
	}
	if b.httpCode != 0 {
		return b.httpCode
	}
	return http.StatusOK
}

func (b *Builder) getContent() interface{} {
	if b.err != nil {
		return b.errorMapper.MapContent(b.err)
	}
	return b.content
}

func SendJSONResponse(res *Response, w http.ResponseWriter) error {
	js, err := json.Marshal(res.Content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return err
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(res.HttpCode)
	w.Write(js)
	return nil
}
