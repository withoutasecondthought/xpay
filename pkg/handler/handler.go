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

	auth := r.Group("/auth")
	{
		auth.POST("/sign-in")
		auth.POST("/sign-up")
	}
	api := r.Group("/api")
	{
		api.GET("/students")
		api.GET("/student")

		api.POST("/new-student")
		api.PUT("/transaction")
		api.DELETE("/remove-student")

	}

	return r
}
