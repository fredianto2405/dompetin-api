package router

import (
	"dompetin-api/internal/auth"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(rg *gin.RouterGroup, handler *auth.Handler) {
	rg.POST("/register", handler.Register)
	rg.POST("/login", handler.Login)
}
