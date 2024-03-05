package user

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *userStorage) Update(ctx context.Context, user *models.User) error {
	return nil
}

func (s *userStorage) Refresh(ctx context.Context, userID, token string) error {
	const op = "userStorage.Refresh"

	query := `UPDATE users SET refresh_token = $1 WHERE id = $2`

	_, err := s.conn.Exec(ctx, query, token, userID)
	if err != nil {
		s.logger.Error("failed to update refresh token", op)
		return models.ErrInternal
	}

	return nil
}
