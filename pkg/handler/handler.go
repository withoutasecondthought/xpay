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

	api := r.Group("/api", h.parseJWT)
	{
		api.GET("/students", h.getStudents)
		api.GET("/student", h.getStudent)

		api.POST("/new-student", h.newStudent)
		api.PUT("/transaction", h.transaction)
		api.DELETE("/remove-student", h.removeStudent)

	}

	return r
}
