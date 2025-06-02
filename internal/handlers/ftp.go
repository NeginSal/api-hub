package handlers

import (
	"github.com/NeginSal/api-hub/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SFTPHandler(c *gin.Context) {
	if err := services.DownloadAndProcessFile(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File downloaded and processed successfully"})
}
