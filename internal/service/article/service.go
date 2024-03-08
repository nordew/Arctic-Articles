package article

import (
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/internal/storage"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"github.com/redis/go-redis/v9"
)

type articleService struct {
	articleStorage storage.ArticleStorage
	redisCl        *redis.Client

	logger logging.Logger
}

func NewArticleService(articleStorage storage.ArticleStorage, redisCl *redis.Client, logger logging.Logger) service.ArticleService {
	return &articleService{
		articleStorage: articleStorage,
		redisCl:        redisCl,

		logger: logger,
	}
}
