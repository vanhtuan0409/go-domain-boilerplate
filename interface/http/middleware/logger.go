package middleware

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/vanhtuan0409/go-domain-boilerplate/infrastructure/logger"
)

type LoggerMiddleware struct{}

func NewLoggerMiddleware() *LoggerMiddleware {
	return &LoggerMiddleware{}
}

func (l *LoggerMiddleware) ServeHTTP(
	rw http.ResponseWriter,
	r *http.Request,
	next http.HandlerFunc,
) {
	next(rw, r)
	res := rw.(negroni.ResponseWriter)
	logger.Logger.Printf(
		"%d %s %s",
		res.Status(),
		r.Method,
		r.URL.Path,
	)
}
