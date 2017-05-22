package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/vanhtuan0409/go-domain-boilerplate/application/accesscontrol"
)

type Handler func(rw http.ResponseWriter, r *http.Request, auth *accesscontrol.AuthInfo)

func adapter(tokenContextKey string, h Handler) func(rw http.ResponseWriter, r *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		authInfo, ok := r.Context().Value(tokenContextKey).(*accesscontrol.AuthInfo)
		if !ok {
			h(rw, r, nil)
			return
		}
		h(rw, r, authInfo)
	}
}

func AdaptHandleFunc(mdw *negroni.Negroni, tokenContextKey string, h Handler) *negroni.Negroni {
	handlers := mdw.Handlers()
	newMdw := negroni.New(handlers...)
	adapted := adapter(tokenContextKey, h)
	newMdw.UseHandlerFunc(adapted)
	return newMdw
}
