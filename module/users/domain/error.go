package domain

import "errors"

var (
	ErrEmailExisted         = errors.New("email is existed")
	ErrInvalidEmailPassword = errors.New("invalid email or password")
	ErrCannotChangeAvatar   = errors.New("cannot change avatar")
	ErrCannotRegister       = errors.New("cannot register")
)
