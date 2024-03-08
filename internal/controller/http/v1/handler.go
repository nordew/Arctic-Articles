package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/nordew/ArcticArticles/pkg/hasher"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"net/http"
)

type Handler struct {
	userService    service.UserService
	articleService service.ArticleService
	feedService    service.FeedService

	auth   auth.Authenticator
	hasher hasher.Hasher
	logger logging.Logger
}

func NewHandler(userService service.UserService, articleService service.ArticleService, feedService service.FeedService, auth auth.Authenticator, hasher hasher.Hasher, logger logging.Logger) *Handler {
	return &Handler{
		userService:    userService,
		articleService: articleService,
		feedService:    feedService,

		auth:   auth,
		hasher: hasher,
		logger: logger,
	}
}

func (h *Handler) Init() *gin.Engine {
	app := gin.New()

	auth := app.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-in", h.signIn)
		auth.GET("/refresh", h.refresh)
	}

	profile := app.Group("/profile")
	profile.Use(h.AuthMiddleware)
	{
		profile.PATCH("/update", h.updateUser)
		profile.DELETE("/", h.deleteUser)
	}

	article := app.Group("/article")
	{
		article.GET("/", h.getArticles)
		article.Use(h.AuthMiddleware)
		article.POST("/", h.createArticle)
		article.DELETE("/:id", h.deleteArticle)
	}

	return app
}

func getTokenFromCtx(c *gin.Context) *auth.ParseTokenClaimsOutput {
	claims, exists := c.Get("tokenClaims")
	if !exists {
		writeErr(c, http.StatusUnauthorized, "claims not found")
		return nil
	}

	tokenClaims, ok := claims.(*auth.ParseTokenClaimsOutput)
	if !ok {
		writeErr(c, http.StatusInternalServerError, "auth error")
		return nil
	}

	return tokenClaims
}

func writeErr(c *gin.Context, statusCode int, msg string) {
	response := gin.H{
		"status":  "error",
		"message": msg,
	}

	c.JSON(statusCode, response)
}

func validationErr(c *gin.Context) {
	response := gin.H{
		"status":  "error",
		"message": "invalid request body",
	}

	c.JSON(http.StatusBadRequest, response)
}

func internalErr(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{})
}
