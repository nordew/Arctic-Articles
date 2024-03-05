package user

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"time"
)

func (s *userService) SignUp(ctx context.Context, user *models.User) (string, string, error) {
	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return "", "", models.ErrInternal
	}

	id := uuid.New().String()

	accessToken, refreshToken, err := s.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: id,
		Role:   0,
	})
	if err != nil {
		return "", "", models.ErrInternal
	}

	userMapped := &models.User{
		ID:           id,
		Name:         user.Name,
		Email:        user.Email,
		Password:     hashedPassword,
		RefreshToken: refreshToken,
		RegisteredAt: time.Now(),
	}

	if err := s.userStorage.Create(ctx, userMapped); err != nil {
		if errors.Is(err, models.ErrEmailAlreadyExists) {
			return "", "", err
		}

		return "", "", models.ErrInternal
	}

	return accessToken, refreshToken, nil
}
