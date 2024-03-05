package user

import (
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/nordew/ArcticArticles/pkg/hasher"
)

type userService struct {
	userStorage storage.UserStorage
	auth        auth.Authenticator
	hasher      hasher.Hasher
}

func NewUserService(userStorage storage.UserStorage, auth auth.Authenticator, hasher hasher.Hasher) service.UserService {
	return &userService{
		userStorage: userStorage,
		auth:        auth,
		hasher:      hasher,
	}
}
