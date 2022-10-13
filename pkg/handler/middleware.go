package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"xpay/pkg/parse"
)

func (h *Handler) parseJWT(c *gin.Context) {
	jwt := c.Request.Header.Get("Authorization")
	errorJSONMessage := gin.H{"message": "invalid token"}
	if jwt == "" {
		c.Status(http.StatusUnauthorized)
		return
	}

	parsedJWT := strings.Split(jwt, " ")
	if len(parsedJWT) != 2 {
		c.JSON(http.StatusUnauthorized, errorJSONMessage)
		return
	}

	if parsedJWT[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, errorJSONMessage)
		return
	}

	id, err := parse.CustomJWTParser(parsedJWT[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": fmt.Sprintf("error occured parsing token: %s", err)})
	}

	c.Set("Authorization", id)
	c.Next()
}
