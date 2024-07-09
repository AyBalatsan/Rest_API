package handler

import (
	"net/http"

	todo "github.com/AyBalatsan/Rest_API"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	// парсит нашу структуру и где binding:"required" не пропустит если поле пустое
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{ // TODO почему тут два {}
		"id": id,
	})

}
func (h *Handler) signIn(c *gin.Context) {

}
