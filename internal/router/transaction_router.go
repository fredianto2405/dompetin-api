package router

import (
	"dompetin-api/internal/transaction"
	"dompetin-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterTransactionRoutes(rg *gin.RouterGroup, handler *transaction.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.Create)
	rg.GET("", handler.History)
	rg.PUT("/:id", handler.Update)
	rg.DELETE("/:id", handler.Delete)
}
