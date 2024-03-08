package articles

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func (s *articleStorage) GetByID(ctx context.Context, articleID string) (*models.Article, error) {
	const op = "articleStorage.GetByID"

	var article models.Article

	query := `
		SELECT id, title, content, author, image_url, date_published 
		FROM articles 
		WHERE id = $1
	`

	err := s.conn.QueryRow(ctx, query, articleID).Scan(
		&article.ArticleID,
		&article.Title,
		&article.Content,
		&article.Author,
		&article.ImageURL,
		&article.DatePublished,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, models.ErrArticleNotFound
		}

		s.logger.Error("failed to scan into article", err.Error(), op)
		return nil, models.ErrInternal
	}

	return &article, nil
}

func (s *articleStorage) GetArticles(ctx context.Context, limit int) ([]models.Article, error) {
	const op = "userStorage.GetArticles"

	query := `SELECT * FROM articles ORDER BY date_published LIMIT $1`

	rows, err := s.conn.Query(ctx, query, limit)
	if err != nil {
		s.logger.Error("failed to execute query", err.Error(), op)
		return nil, models.ErrInternal
	}
	defer rows.Close()

	var articles []models.Article

	for rows.Next() {
		var article models.Article

		if err := rows.Scan(
			&article.ArticleID,
			&article.Title,
			&article.Content,
			&article.Author,
			&article.ImageURL,
			&article.DatePublished,
		); err != nil {
			s.logger.Error("failed to scan row", err.Error(), op)

			return nil, models.ErrInternal
		}

		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		s.logger.Error("error after reading all rows", err.Error(), op)

		return nil, models.ErrInternal
	}

	return articles, nil
}
