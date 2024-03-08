package v1

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/nordew/ArcticArticles/internal/controller/dto"
	"github.com/nordew/ArcticArticles/internal/domain/models"
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

func (h *Handler) updateUser(c *gin.Context) {
	token := getTokenFromCtx(c)

	var updateUserDTO dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&updateUserDTO); err != nil {
		validationErr(c)
		return
	}

	user, err := h.userService.GetByID(context.Background(), token.Sub)
	if err != nil {
		internalErr(c)
		return
	}

	if updateUserDTO.Name != "" {
		if err := models.ValidateUserName(updateUserDTO.Name); err != nil {
			writeErr(c, http.StatusBadRequest, err.Error())
			return
		}

		user.Name = updateUserDTO.Name
	}

	if updateUserDTO.OldPassword != "" && updateUserDTO.NewPassword != "" {
		if err := models.ValidateUserPassword(updateUserDTO.NewPassword); err != nil {
			writeErr(c, http.StatusBadRequest, err.Error())
			return
		}

		if err := h.hasher.Compare(user.Password, updateUserDTO.OldPassword); err != nil {
			writeErr(c, http.StatusBadRequest, "Invalid old password")
			return
		}

		hashedPassword, err := h.hasher.Hash(updateUserDTO.NewPassword)
		if err != nil {
			writeErr(c, http.StatusInternalServerError, err.Error())
			return
		}

		user.Password = hashedPassword
	}

	if err := h.userService.Update(context.Background(), user); err != nil {
		writeErr(c, http.StatusInternalServerError, err.Error())
		return
	}
}
