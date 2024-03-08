package v1

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nordew/ArcticArticles/internal/controller/dao"
	"github.com/nordew/ArcticArticles/internal/controller/dto"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"log"
	"net/http"
	"time"
)

func (h *Handler) createArticle(c *gin.Context) {
	var articleDTO dto.CreateArticleDTO

	if err := c.ShouldBindJSON(&articleDTO); err != nil {
		h.logger.Error("Error binding JSON", err)
		internalErr(c)
		return
	}

	convertedArticle := dao.ToArticleFromDTO(&articleDTO)

	if err := convertedArticle.Validate(); err != nil {
		h.logger.Error("Validation error", err)
		validationErr(c)
		return
	}

	token := getTokenFromCtx(c)

	user, err := h.userService.GetByID(context.Background(), token.Sub)
	if err != nil {
		if errors.Is(err, models.ErrWrongEmailOrPassword) {
			h.logger.Error("Invalid token", err)

			writeErr(c, http.StatusUnauthorized, "invalid token")
			return
		}

		h.logger.Error("User service error", err.Error())

		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	articleParsed := &models.Article{
		ArticleID:     uuid.New().String(),
		Title:         convertedArticle.Title,
		Content:       convertedArticle.Content,
		Author:        user.Name,
		DatePublished: time.Now(),
		//ImageURL:      "",
	}

	if err := h.articleService.Create(context.Background(), articleParsed); err != nil {
		h.logger.Error("Article service error", err)

		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}

func (h *Handler) deleteArticle(c *gin.Context) {
	token := getTokenFromCtx(c)

	id := c.Param("id")

	log.Println(id)

	article, err := h.articleService.GetByID(context.Background(), id)
	if err != nil {
		if errors.Is(err, models.ErrArticleNotFound) {
			writeErr(c, http.StatusNotFound, err.Error())
			return
		}

		writeErr(c, http.StatusInternalServerError, err.Error())
	}

	user, err := h.userService.GetByID(context.Background(), token.Sub)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			writeErr(c, http.StatusNotFound, err.Error())
			return
		}

		writeErr(c, http.StatusInternalServerError, err.Error())
	}

	if article.Author != user.Name {
		writeErr(c, http.StatusUnauthorized, "access denied")
	}

	if err := h.articleService.Delete(context.Background(), id); err != nil {
		writeErr(c, http.StatusUnauthorized, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) getArticles(c *gin.Context) {
	articles, err := h.feedService.GetArticles(context.Background())
	if err != nil {
		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, articles)
}
