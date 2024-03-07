package models

import (
	"github.com/go-playground/validator/v10"
	"time"
	"unicode/utf8"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	ID           string
	Email        string
	Name         string
	Password     string
	RefreshToken string
	Role         int
	RegisteredAt time.Time
}

type SignUpInput struct {
	Name     string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type SignInInput struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

func (i *SignUpInput) Validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	if utf8.RuneCountInString(i.Name) > 32 || utf8.RuneCountInString(i.Name) < 2 {
		return ErrWrongNameLength
	}

	return nil
}

func (i *SignInInput) Validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}
