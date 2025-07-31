package router

import (
	"dompetin-api/internal/auth"
	"dompetin-api/internal/user"
	"dompetin-api/pkg/errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	errors.InitValidator()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           15 * time.Minute,
	}))

	r.Use(errors.ErrorHandler())

	// auth routes
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	authService := auth.NewService(userService)
	authHandler := auth.NewHandler(authService, userService)
	authGroup := r.Group("/api/v1/auth")
	RegisterAuthRoutes(authGroup, authHandler)

	return r
}
