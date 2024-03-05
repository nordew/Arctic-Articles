package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nordew/ArcticArticles/internal/domain/models"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr(c)
		return
	}

	if err := input.Validate(); err != nil {
		validationErr(c)
		return
	}

	user := &models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}

	accessToken, refreshToken, err := h.userService.SignUp(context.Background(), user)
	if err != nil {
		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationErr(c)
		return
	}

	if err := input.Validate(); err != nil {
		validationErr(c)
		return
	}

	accessToken, refreshToken, err := h.userService.SignIn(context.Background(), input.Email, input.Password)
	if err != nil {
		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *Handler) refresh(c *gin.Context) {
	token := c.GetHeader("refresh_token")

	accessToken, refreshToken, err := h.userService.Refresh(context.Background(), token)
	if err != nil {
		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}

	c.JSON(http.StatusOK, response)
}