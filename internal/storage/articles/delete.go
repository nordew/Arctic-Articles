package articles

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *articleStorage) Delete(ctx context.Context, articleID string) error {
	const op = "articleStorage.Delete"

	query := `DELETE FROM articles WHERE id = $1`

	result, err := s.conn.Exec(ctx, query, articleID)
	if err != nil {
		s.logger.Error("failed to delete article", err.Error(), op)

		return err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return models.ErrArticleNotFound
	}

	return nil
}
