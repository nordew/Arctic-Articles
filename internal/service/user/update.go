package user

import (
	"context"
	"github.com/nordew/ArcticArticles/pkg/auth"
)

func (s *userService) Refresh(ctx context.Context, token string) (string, string, error) {
	user, err := s.userStorage.GetByToken(ctx, token)
	if err != nil {
		return "", "", err
	}

	accessToken, refreshToken, err := s.auth.GenerateTokens(&auth.GenerateTokenClaimsOptions{
		UserId: user.ID,
		Role:   user.Role,
	})
	if err != nil {
		return "", "", err
	}

	if err := s.userStorage.Refresh(ctx, user.ID, refreshToken); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
