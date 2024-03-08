package article

import (
	"context"
	"encoding/json"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"time"
)

func (s *articleService) Create(ctx context.Context, article *models.Article) error {
	marshalledArticle, err := json.Marshal(article)
	if err != nil {
		s.logger.Error("failed to marshal article", err.Error())

		return models.ErrInternal
	}

	if err := s.redisCl.Set(ctx, article.ArticleID, marshalledArticle, time.Hour*24).Err(); err != nil {
		s.logger.Error("failed to set into redis", err.Error())

		return models.ErrInternal
	}

	if err := s.articleStorage.Create(ctx, article); err != nil {
		return models.ErrInternal
	}

	return nil
}
