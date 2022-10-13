package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xpay"
)

func (h *Handler) signIn(c *gin.Context) {
	var teacher xpay.Teacher

	err := c.BindJSON(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid request: %s", err)})
		return
	}

	jwt, err := h.services.Auth.SignIn(teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("invalid request: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func (h *Handler) signUp(c *gin.Context) {
	var teacher xpay.Teacher

	err := c.BindJSON(&teacher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("invalid request: %s", err)})
		return
	}

	jwt, err := h.services.Auth.SignUp(teacher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("invalid request: %s", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
