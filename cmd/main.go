package main

import (
	"log"
	"net/http"

	"github.com/NeginSal/api-hub/internal/config"
	"github.com/NeginSal/api-hub/internal/middleware"

	"github.com/NeginSal/api-hub/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	// r := gin.Default()
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware())
	r.Use(middleware.RateLimitMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.POST("/login", handlers.LoginHandler)

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to protected route!"})
	})

	// start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server : %v", err)
	}

}
