package feed

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *feedService) GetArticles(ctx context.Context) ([]models.Article, error) {
	return s.articleStorage.GetArticles(ctx, s.limit)
}
