package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) deleteUser(c *gin.Context) {
	token := getTokenFromCtx(c)

	if err := h.userService.Delete(context.Background(), token.Sub); err != nil {
		h.logger.Error("User service error", err.Error())
		writeErr(c, http.StatusInternalServerError, "failed to delete user")
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
