package http

import (
    "net/http"
	"encoding/json"
)

type Response struct {
    HttpCode int
    Content interface{}
    Error error
}

func (r *Response) GetHttpCode() int {
    if r.HttpCode != 0 {
        return r.HttpCode
    }
    return GetStatusCodeFromError(r.Error)
}

func (r *Response) GetJSON() ([]byte, error) {
    return json.Marshal(r.Content)
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

func SendJSONResponse(res Response, w http.ResponseWriter) error {
    js, err := res.GetJSON()
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write("")
        return err
    }
    w.WriteHeader(res.GetHttpCode())
    w.Write(js)
    return nil
}
