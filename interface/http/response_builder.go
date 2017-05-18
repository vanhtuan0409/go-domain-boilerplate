package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	HttpCode int
	Content  interface{}
	Error    error
}

func (r *Response) GetHttpCode() int {
	if r.Error != nil {
		return GetStatusCodeFromError(r.Error)
	}
	return r.HttpCode
}

func (r *Response) GetJSON() ([]byte, error) {
	if r.Error != nil {
		return []byte(r.Error.Error()), nil
	}
	return json.Marshal(r.Content)
}

type Builder struct {
	httpCode int
	content  interface{}
	err      error
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

func (b *Builder) Error(err error) *Builder {
	b.err = err
	return b
}

func (b *Builder) Build() *Response {
	response := Response{}
	response.HttpCode = b.httpCode
	response.Content = b.content
	response.Error = b.err
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
