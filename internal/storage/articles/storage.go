package articles

import (
	"github.com/jackc/pgx/v5"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/logging"
)

type articleStorage struct {
	conn *pgx.Conn

	logger logging.Logger
}

func NewArticleStorage(conn *pgx.Conn, logger logging.Logger) storage.ArticleStorage {
	return &articleStorage{
		conn:   conn,
		logger: logger,
	}
}
