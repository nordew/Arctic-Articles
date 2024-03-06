package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nordew/ArcticArticles/pkg/auth"
	"net/http"
)

func (h *Handler) delete(c *gin.Context) {
	claims, exists := c.Get("tokenClaims")
	if !exists {
		writeErr(c, http.StatusUnauthorized, "claims not found")
		return
	}

	tokenClaims, ok := claims.(*auth.ParseTokenClaimsOutput)
	if !ok {
		writeErr(c, http.StatusInternalServerError, "auth error")
		return
	}

	if err := h.userService.Delete(context.Background(), tokenClaims.Sub); err != nil {
		writeErr(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
