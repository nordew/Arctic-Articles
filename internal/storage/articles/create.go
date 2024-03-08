package articles

import (
	"context"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *articleStorage) Create(ctx context.Context, article *models.Article) error {
	const op = "articleStorage.Create"

	query := `INSERT INTO articles (id, title, content, author, image_url, date_published) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := s.conn.Exec(ctx, query, article.ArticleID, article.Title, article.Content, article.Author, article.ImageURL, article.DatePublished)
	if err != nil {
		s.logger.Error("failed to create article", err.Error())

		return err
	}

	return nil
}
