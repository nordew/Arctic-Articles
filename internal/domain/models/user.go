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
	ID           string    `json:"id"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	RefreshToken string    `json:"refresh_token"`
	Role         int       `json:"role"`
	RegisteredAt time.Time `json:"registered_at"`
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

	if err := ValidateUserName(i.Name); err != nil {
		return err
	}

	if err := ValidateUserPassword(i.Password); err != nil {
		return err
	}

	return nil
}

func (i *SignInInput) Validate() error {
	if err := validate.Struct(i); err != nil {
		return err
	}

	return nil
}

func ValidateUserName(name string) error {
	if utf8.RuneCountInString(name) > 32 || utf8.RuneCountInString(name) < 2 {
		return ErrWrongNameLength
	}

	return nil
}

func ValidateUserPassword(password string) error {
	if utf8.RuneCountInString(password) < 8 {
		return ErrWrongPasswordLength
	}

	return nil
}
