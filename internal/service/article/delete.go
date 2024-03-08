package article

import (
	"context"
)

func (s *articleService) Delete(ctx context.Context, articleID string) error {
	existsInRedis, err := s.redisCl.Exists(ctx, articleID).Result()
	if err != nil {
		s.logger.Error("failed to check article existence in Redis", err.Error())
	} else if existsInRedis == 1 {
		if err := s.redisCl.Del(ctx, articleID).Err(); err != nil {
			s.logger.Error("failed to delete article from Redis", err.Error())
		}
	}

	return s.articleStorage.Delete(ctx, articleID)
}
