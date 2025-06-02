// @title API Hub
// @version 1.0
// @description Microservice Gateway to LMS, HMS, BMS
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"log"
	"net/http"

	"github.com/NeginSal/api-hub/internal/config"
	"github.com/NeginSal/api-hub/internal/middleware"

	"github.com/NeginSal/api-hub/internal/handlers"
	"github.com/gin-gonic/gin"

	_ "github.com/NeginSal/api-hub/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadEnv()
	config.ConnectMongo()

	r := gin.New()
	r.Use(gin.Recovery(), middleware.LoggerMiddleware(), middleware.RateLimitMiddleware())

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "OK"})
	})

	r.POST("/login", handlers.LoginHandler)

	// Swagger route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	protected := r.Group("/protected")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to protected route!"})
		})

		protected.GET("/lms", handlers.FetchLMSHandler)
		protected.GET("/sftp", handlers.SFTPHandler)
	}

	// start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server : %v", err)
	}

}
