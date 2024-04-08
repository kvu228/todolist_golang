package common

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrBadRequest     = errors.New("bad request")
)
