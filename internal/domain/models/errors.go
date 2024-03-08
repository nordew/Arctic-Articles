package models

import "errors"

var (
	ErrInternal             = errors.New("internal server error")
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
	ErrUserNotFound         = errors.New("user not found")
	ErrArticleNotFound      = errors.New("article not found")
	ErrInvalidToken         = errors.New("invalid token")
	ErrWrongNameLength      = errors.New("name length is more than 32 or less than 8 symbols")
	ErrWrongPasswordLength  = errors.New("password length must be more than 8 symbols")
	ErrWrongTitleLength     = errors.New("title length is more than 70 symbols")
)
