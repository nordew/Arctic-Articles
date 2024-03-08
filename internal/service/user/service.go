package user

import (
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/nordew/ArcticArticles/pkg/hasher"
	"github.com/redis/go-redis/v9"
)

type userService struct {
	userStorage storage.UserStorage
	redisCl     *redis.Client

	auth   auth.Authenticator
	hasher hasher.Hasher
}

func NewUserService(userStorage storage.UserStorage, redisCl *redis.Client, auth auth.Authenticator, hasher hasher.Hasher) service.UserService {
	return &userService{
		userStorage: userStorage,
		redisCl:     redisCl,

		auth:   auth,
		hasher: hasher,
	}
}
