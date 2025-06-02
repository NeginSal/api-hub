package handlers

import (
	"net/http"

	"github.com/NeginSal/api-hub/internal/services"
	"github.com/gin-gonic/gin"
)

func FetchLMSHandler(c *gin.Context) {
	if err := services.FetchAndStoreLMSData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "LMS data stored successfully"})
}
