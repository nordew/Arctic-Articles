package feed

import (
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"github.com/redis/go-redis/v9"
)

type feedService struct {
	redisCl        *redis.Client
	articleStorage storage.ArticleStorage

	logger logging.Logger

	limit int
}

func NewFeedService(redisCl *redis.Client, articleStorage storage.ArticleStorage, logger logging.Logger, limit int) service.FeedService {
	return &feedService{
		redisCl:        redisCl,
		articleStorage: articleStorage,

		logger: logger,

		limit: limit,
	}
}
