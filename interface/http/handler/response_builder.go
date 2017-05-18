package handler

import (
	"encoding/json"
	"net/http"
)

type IErrorMapper interface {
	MapHttpCode(err error) int
	MapContent(err error) interface{}
}

type Response struct {
	httpCode    int
	content     interface{}
	err         error
	errorMapper IErrorMapper
}

func (r *Response) GetHttpCode() int {
	if r.err != nil {
		return r.errorMapper.MapHttpCode(r.err)
	}
	if r.httpCode != 0 {
		return r.httpCode
	}
	return http.StatusOK
}

func (r *Response) GetJSON() ([]byte, error) {
	if r.err != nil {
		errorContent := r.errorMapper.MapContent(r.err)
		return json.Marshal(errorContent)
	}
	return json.Marshal(r.content)
}

type Builder struct {
	httpCode    int
	content     interface{}
	err         error
	errorMapper IErrorMapper
}

func ReponseBuilder(mapper IErrorMapper) *Builder {
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
	response.httpCode = b.httpCode
	response.content = b.content
	response.err = b.err
	response.errorMapper = b.errorMapper
	return &response
}

func SendJSONResponse(res *Response, w http.ResponseWriter) error {
	js, err := res.GetJSON()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	w.WriteHeader(res.GetHttpCode())
	w.Write(js)
	return nil
}
