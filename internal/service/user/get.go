package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/redis/go-redis/v9"
	"time"
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

func (s *userService) GetByID(ctx context.Context, userID string) (*models.User, error) {
	cachedUser, err := s.redisCl.Get(ctx, userID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			user, err := s.userStorage.GetByID(ctx, userID)
			if err != nil {
				if errors.Is(err, models.ErrUserNotFound) {
					return nil, err
				}

				return nil, models.ErrInternal
			}

			marshalledUser, err := json.Marshal(user)
			if err != nil {
				return nil, models.ErrInternal
			}

			if err := s.redisCl.Set(ctx, userID, marshalledUser, time.Hour*3).Err(); err != nil {
				return nil, models.ErrInternal
			}

			return user, nil
		}

		return nil, models.ErrInternal
	}

	var user models.User
	if err := json.Unmarshal([]byte(cachedUser), &user); err != nil {
		return nil, models.ErrInternal
	}

	return &user, nil
}
