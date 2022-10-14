package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xpay"
)

func (h *Handler) getStudents(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"id is": c.MustGet("Authorization")})
}

func (h *Handler) getStudent(c *gin.Context) {

}

func (h *Handler) newStudent(c *gin.Context) {
	var student xpay.Student

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("error occurred binding JSON: %s", err)})
		return
	}

	var teacherId = c.GetInt("Authorization")
	if teacherId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization error"})
		return
	}

	student.TeacherID = teacherId

}

func (h *Handler) transaction(c *gin.Context) {

}

func (h *Handler) removeStudent(c *gin.Context) {

}
