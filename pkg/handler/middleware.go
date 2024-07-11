package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorzationHeader = "Authorization"
	userCtx            = "userId"
)

func (h *Handler) userIndentily(c *gin.Context) {
	header := c.GetHeader(authorzationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
	}
	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	idUser, ok := c.Get(userCtx)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}
	id, ok := idUser.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id is of invalid type")
	}
	return id, nil
}
