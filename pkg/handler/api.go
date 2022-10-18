package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getStudents(c *gin.Context) {
	teacherId, exist := c.Get("Authorization")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"id is": teacherId})
}

func (h *Handler) getStudent(c *gin.Context) {

}

func (h *Handler) newStudent(c *gin.Context) {
	teacherId, exist := c.Get("Authorization")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"id is": teacherId})
}

func (h *Handler) transaction(c *gin.Context) {
	teacherId, exist := c.Get("Authorization")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"id is": teacherId})
}

func (h *Handler) removeStudent(c *gin.Context) {
	teacherId, exist := c.Get("Authorization")
	if !exist {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"id is": teacherId})
}
