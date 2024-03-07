package dao

import (
	"github.com/nordew/ArcticArticles/internal/controller/dto"
	"github.com/nordew/ArcticArticles/internal/domain/models"
)

func ToArticleFromDTO(articleDTO *dto.CreateArticleDTO) *models.Article {
	return &models.Article{
		Title:   articleDTO.Title,
		Content: articleDTO.Content,
	}
}
