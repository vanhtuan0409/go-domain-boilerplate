package http

import (
	"errors"
)

var (
	ErrorParseCheckInRequest = errors.New("Parse Check in request failed")
	ErrorInvalidToken        = errors.New("Invalid Token")
)
