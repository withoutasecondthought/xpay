package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getStudents(c *gin.Context) {

}

func (h *Handler) getStudent(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id is": c.Request.Header.Get("Authorization")})
}

func (h *Handler) newStudent(c *gin.Context) {

}

func (h *Handler) transaction(c *gin.Context) {

}

func (h *Handler) removeStudent(c *gin.Context) {

}
