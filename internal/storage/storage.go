package storage

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

// UserStorage provides methods for user data manipulation.
type UserStorage interface {
	// Create creates a new user in the database.
	Create(ctx context.Context, user *models.User) error

	// GetByID retrieves a user from the database by their ID.
	GetByID(ctx context.Context, userID string) (*models.User, error)

	// GetByEmail retrieves a user from the database by their email address.
	GetByEmail(ctx context.Context, email string) (*models.User, error)

	// GetByToken retrieves a user from the database by their refresh token.
	GetByToken(ctx context.Context, token string) (*models.User, error)

	// Update updates the information of an existing user in the database.
	Update(ctx context.Context, user *models.User) error

	// Refresh updates the refresh token for a user in the database.
	Refresh(ctx context.Context, userID, token string) error
}
