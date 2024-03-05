package user

import (
	"context"
	"errors"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"github.com/nordew/ArcticArticles/pkg/auth"
)

func (s *userService) SignIn(ctx context.Context, email, password string) (string, string, error) {
	user, err := s.userStorage.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, models.ErrWrongEmailOrPassword) {
			return "", "", err
		}

		return "", "", models.ErrInternal
	}

	if err := s.hasher.Compare(user.Password, password); err != nil {
		return "", "", models.ErrWrongEmailOrPassword
	}

	accessToken, refreshToken, err := s.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.ID,
		Role:   user.Role,
	})
	if err != nil {
		return "", "", models.ErrInternal
	}

	if err := s.userStorage.Refresh(ctx, user.ID, refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
