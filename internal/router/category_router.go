package router

import (
	"dompetin-api/internal/category"
	"dompetin-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, handler *category.Handler) {
	rg.Use(middleware.JWTAuthMiddleware())
	rg.POST("", handler.Create)
	rg.GET("", handler.GetByUserID)
	rg.PUT("/:id", handler.Update)
	rg.DELETE("/:id", handler.Delete)
}
