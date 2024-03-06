package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AuthMiddleware(c *gin.Context) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	claims, err := h.auth.ParseToken(accessToken)
	if err != nil {
		h.logger.Info("parseToken", err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		c.Abort()
		return
	}

	c.Set("tokenClaims", claims)

	c.Next()
}
