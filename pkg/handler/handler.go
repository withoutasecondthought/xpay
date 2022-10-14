package handler

import (
	"github.com/gin-gonic/gin"
	"xpay/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		services: s,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/sign-in", h.signIn)
	r.POST("/sign-up", h.signUp)
	r.GET("/student", h.getStudent)

	api := r.Group("/api", h.parseJWT)
	{
		api.GET("/students", h.getStudents)
		api.POST("/new-student", h.newStudent)
		api.POST("/transaction", h.transaction)
		api.DELETE("/remove-student", h.removeStudent)

	}

	return r
}
