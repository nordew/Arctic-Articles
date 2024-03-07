package models

import "errors"

var (
	ErrInternal             = errors.New("internal server error")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrInvalidToken         = errors.New("invalid token")
	ErrWrongNameLength      = errors.New("name length is more than 32 symbols")
)
