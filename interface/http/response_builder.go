package http

import (
    "net/http"
)

type Response struct {
    HttpCode int
    Content interface{}
    Error error
}

type Builder struct {
    httpCode int
    content interface{}
    err error
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

func SendResponse(res Response, w http.ResponseWriter) error {
    return nil
}
