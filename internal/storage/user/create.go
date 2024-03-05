package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *userStorage) Create(ctx context.Context, user *models.User) error {
	const op = "userStorage.Create"

	query := `
		INSERT INTO users (id, name, email, password_hash, refresh_token, registered_at)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := s.conn.Exec(ctx, query, user.ID, user.Name, user.Email, user.Password, user.RefreshToken, user.RegisteredAt)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == "23505" {
			return models.ErrEmailAlreadyExists
		}

		s.logger.Error("failed to create user: %s", err.Error(), err, op)
		return err
	}

	return nil
}
