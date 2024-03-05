package user

import (
	"github.com/jackc/pgx/v5"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/logging"
)

type userStorage struct {
	conn *pgx.Conn

	logger logging.Logger
}

func NewUserStorage(conn *pgx.Conn, logger logging.Logger) storage.UserStorage {
	return &userStorage{
		conn:   conn,
		logger: logger,
	}
}
