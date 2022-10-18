package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"xpay"
)

func (h *Handler) getStudents(c *gin.Context) {
	teacherId := c.GetInt("Authorization")
	if teacherId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
	}

	students, err := h.services.Student.GetStudents(teacherId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error getting students: %s", err)})
		return
	}
	c.JSON(http.StatusOK, students)
}

func (h *Handler) getStudent(c *gin.Context) {
	strId := c.Param("id")
	if strId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no id"})
		return
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	student, err := h.services.Student.GetStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("cant get student: %s", err)})
		return
	}

	c.JSON(http.StatusOK, student)
}

func (h *Handler) newStudent(c *gin.Context) {
	var student xpay.Student

	teacherId := c.GetInt("Authorization")
	if teacherId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
		return
	}

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error occuring json: %s", err))
		return
	}

	student.TeacherId = teacherId

	err = h.services.Student.NewStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error getting student: %s", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) transaction(c *gin.Context) {
	var transaction xpay.Transaction
	teacherId := c.GetInt("Authorization")
	if teacherId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
		return
	}

	err := c.BindJSON(&transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error occuring json: %s", err))
		return
	}

	transaction.TeacherId = teacherId

	err = h.services.Student.Transaction(transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error occured while transaction: %s", err)})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) removeStudent(c *gin.Context) {
	var student xpay.Student
	teacherId := c.GetInt("Authorization")
	if teacherId == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "error Unauthorized"})
		return
	}

	err := c.BindJSON(&student)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("error occuring json: %s", err))
		return
	}

	student.TeacherId = teacherId

	err = h.services.Student.DeleteStudent(student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("error occured while delete student: %s", err)})
		return
	}
	c.Status(http.StatusOK)
}
