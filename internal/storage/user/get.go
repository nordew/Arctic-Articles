package user

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *userStorage) GetByID(ctx context.Context, userID string) (*models.User, error) {
	const op = "userStorage.GetByID"

	sqlQuery := `SELECT id, name, email, password_hash, role, refresh_token, registered_at FROM users WHERE id = $1`

	var user models.User

	err := s.conn.QueryRow(ctx, sqlQuery, userID).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.RefreshToken,
		&user.RegisteredAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}

		s.logger.Error("failed to get user", err.Error(), op)
		return nil, err
	}

	return &user, nil
}

func (s *userStorage) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	const op = "userStorage.GetByEmail"

	sqlQuery := `SELECT id, name, email, password_hash, role, refresh_token, registered_at FROM users WHERE email = $1`

	var user models.User

	err := s.conn.QueryRow(ctx, sqlQuery, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.RefreshToken,
		&user.RegisteredAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrUserNotFound
		}

		s.logger.Error("failed to get user", err.Error(), op)
		return nil, err
	}

	return &user, nil
}

func (s *userStorage) GetByToken(ctx context.Context, token string) (*models.User, error) {
	const op = "userStorage.GetByToken"

	query := `SELECT * FROM users WHERE refresh_token = $1`

	var user models.User

	err := s.conn.QueryRow(ctx, query, token).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.RefreshToken,
		&user.Role,
		&user.RegisteredAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			s.logger.Error("user not found", err.Error(), op)
			return nil, models.ErrInvalidToken
		}

		s.logger.Error("failed to get user", err.Error(), op)
		return nil, models.ErrInternal
	}

	return &user, nil
}
