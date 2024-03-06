package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/nordew/ArcticArticles/internal/service"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"github.com/nordew/ArcticArticles/pkg/logging"
	"net/http"
)

type Handler struct {
	userService service.UserService
	auth        auth.Authenticator
	logger      logging.Logger
}

func NewHandler(userService service.UserService, auth auth.Authenticator, logger logging.Logger) *Handler {
	return &Handler{
		userService: userService,
		auth:        auth,
		logger:      logger,
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
		profile.DELETE("/", h.delete)
	}

	return app
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
