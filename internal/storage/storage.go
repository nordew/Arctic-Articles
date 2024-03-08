package storage

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

// UserStorage provides methods for managing user data in the database.
type UserStorage interface {
	// Create inserts a new user into the database.
	// It takes a context and a *models.User object representing the user to be created.
	// Returns an error if the creation operation fails, otherwise returns nil.
	Create(ctx context.Context, user *models.User) error

	// GetByID retrieves a user from the database based on the provided user ID.
	// It takes a context and a string representing the user ID.
	// Returns a *models.User representing the retrieved user if found, and an error.
	// If the user with the given ID does not exist, it returns nil for the user and an error.
	GetByID(ctx context.Context, userID string) (*models.User, error)

	// GetByEmail retrieves a user from the database based on the provided email address.
	// It takes a context and a string representing the email address.
	// Returns a *models.User representing the retrieved user if found, and an error.
	// If the user with the given email address does not exist, it returns nil for the user and an error.
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// GetByToken retrieves a user from the database based on the provided refresh token.
	// It takes a context and a string representing the refresh token.
	// Returns a *models.User representing the retrieved user if found, and an error.
	// If the user with the given refresh token does not exist, it returns nil for the user and an error.
	GetByToken(ctx context.Context, token string) (*models.User, error)

	// Update updates the information of an existing user in the database.
	// It takes a context and a *models.User object containing the updated user information.
	// Returns an error if the update operation fails, otherwise returns nil.
	Update(ctx context.Context, user *models.User) error

	// Refresh updates the refresh token for a user in the database.
	// It takes a context, a string representing the user ID, and a string representing the new refresh token.
	// Returns an error if the refresh operation fails, otherwise returns nil.
	Refresh(ctx context.Context, userID, token string) error

	// Delete completely deletes a user from the database.
	// It takes a context and a string representing the user ID.
	// Returns an error if the deletion operation fails, otherwise returns nil.
	Delete(ctx context.Context, userID string) error
}

// ArticleStorage defines methods for managing articles in the data store.
type ArticleStorage interface {
	// Create inserts a new article into the data store.
	// It takes a context and a *models.Article object representing the article to be created.
	// Returns an error if the creation operation fails, otherwise returns nil.
	Create(ctx context.Context, article *models.Article) error

	// GetByID retrieves an article from the data store based on the provided article ID.
	// It takes a context and a string representing the article ID.
	// Returns a *models.Article representing the retrieved article if found, and an error.
	// If the article with the given ID does not exist, it returns nil for the article and an error.
	GetByID(ctx context.Context, articleID string) (*models.Article, error)

	// GetArticles retrieves a list of articles from the data store with the provided limit.
	// It takes a context and an integer representing the maximum number of articles to retrieve.
	// Returns a slice of models.Article representing the retrieved articles and an error.
	GetArticles(ctx context.Context, limit int) ([]models.Article, error)

	// Delete deletes an article with the provided article ID from the data store.
	// It takes a context and a string representing the article ID.
	// Returns an error if the deletion operation fails, otherwise returns nil.
	Delete(ctx context.Context, articleID string) error
}
