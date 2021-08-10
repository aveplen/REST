package constants

import "errors"

var (
	ErrAlreadyExists = errors.New("user with given email already exists")
)
