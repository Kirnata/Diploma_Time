package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New() // gin without default middleware
	auth := router.Group("/auth")
	{
		auth.POST("/sign_up", h.signUp)
		auth.POST("/sign_in", h.signIn)
	}
	return router
}
