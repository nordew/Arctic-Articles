package article

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"github.com/redis/go-redis/v9"
	"time"
)

func (s *articleService) GetByID(ctx context.Context, articleID string) (*models.Article, error) {
	cachedArticle, err := s.redisCl.Get(ctx, articleID).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			article, err := s.articleStorage.GetByID(ctx, articleID)
			if err != nil {
				if errors.Is(err, models.ErrArticleNotFound) {
					return nil, err
				}
				return nil, models.ErrInternal
			}

			marshalledArticle, err := json.Marshal(article)
			if err != nil {
				return nil, models.ErrInternal
			}

			if err := s.redisCl.Set(ctx, articleID, marshalledArticle, time.Hour*24).Err(); err != nil {
				return nil, models.ErrInternal
			}

			return article, nil
		}

		return nil, models.ErrInternal
	}

	var article models.Article
	if err := json.Unmarshal([]byte(cachedArticle), &article); err != nil {
		return nil, models.ErrInternal
	}

	return &article, nil
}
