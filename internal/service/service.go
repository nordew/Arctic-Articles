package service

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

// UserService defines methods for user authentication and management.
type UserService interface {
	// SignUp registers a new user with the provided details and returns access and refresh tokens upon successful registration.
	// Returns accessToken, refreshToken and an error. If an error occurs, both tokens will be empty strings.
	SignUp(ctx context.Context, user *models.User) (string, string, error)

	// SignIn authenticates a user with the provided email and password and returns access and refresh tokens upon successful authentication.
	// Returns accessToken, refreshToken and an error. If an error occurs, both tokens will be empty strings.
	SignIn(ctx context.Context, email, password string) (string, string, error)

	// Refresh generates new access and refresh tokens based on the provided refresh token.
	// Returns new accessToken, refreshToken and an error. If an error occurs, both tokens will be empty strings.
	Refresh(ctx context.Context, token string) (string, string, error)
}
