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

	// GetByID retrieves a user from the data store based on the provided user ID.
	// It returns a *models.User representing the retrieved user if found, and nil for the error.
	// If the user with the given userID does not exist, it returns nil for the user and an error.
	GetByID(ctx context.Context, userID string) (*models.User, error)

	// SignIn authenticates a user with the provided email and password and returns access and refresh tokens upon successful authentication.
	// Returns accessToken, refreshToken and an error. If an error occurs, both tokens will be empty strings.
	SignIn(ctx context.Context, email, password string) (string, string, error)

	// Refresh generates new access and refresh tokens based on the provided refresh token.
	// Returns new accessToken, refreshToken and an error. If an error occurs, both tokens will be empty strings.
	Refresh(ctx context.Context, token string) (string, string, error)

	// Update updates the user information in the data store based on the provided user object.
	// It takes a context and a *models.User object containing the updated user information.
	Update(ctx context.Context, user *models.User) error

	// Delete deletes a user with the provided userID from the system.
	// It removes all user-related data associated with the given userID.
	Delete(ctx context.Context, userID string) error
}

// ArticleService defines methods for article management.
type ArticleService interface {
	// Create adds a new article to the system.
	// It takes a context and a *models.Article object representing the article to be created.
	// Returns an error if the creation operation fails, otherwise returns nil.
	Create(ctx context.Context, article *models.Article) error

	// GetByID retrieves an article from the system based on the provided article ID.
	// It takes a context and a string representing the article ID.
	// Returns a *models.Article representing the retrieved article if found, and an error.
	// If the article with the given ID does not exist, it returns nil for the article and an error.
	GetByID(ctx context.Context, articleID string) (*models.Article, error)

	// Delete removes an article from the system based on the provided article ID.
	// It takes a context and a string representing the article ID.
	// Returns an error if the deletion operation fails, otherwise returns nil.
	Delete(ctx context.Context, articleID string) error
}

// FeedService defines methods for retrieving articles for the feed.
type FeedService interface {
	// GetArticles retrieves a list of articles for the feed.
	// It takes a context.
	// Returns a slice of models.Article representing the articles in the feed and an error.
	GetArticles(ctx context.Context) ([]models.Article, error)
}
