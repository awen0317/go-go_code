package model

import (
	"errors"
)

var (
	ErrUserNotExit   = errors.New("user not exit")
	ErrInvalidPasswd = errors.New("Passwd or use not right")
	ErrInvalidParams = errors.New("Invalid params")
	ErrUserHasExit   = errors.New("user has exit")

)
