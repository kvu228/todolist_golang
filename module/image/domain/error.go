package domain

import "errors"

var (
	ErrCannotUploadImage = errors.New("cannot upload image")
	ErrCannotFileImage   = errors.New("cannot find image")
)
