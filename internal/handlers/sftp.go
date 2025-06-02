package handlers

import (
	"github.com/NeginSal/api-hub/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SFTPHandler godoc
// @Summary Fetch file from SFTP and print contents
// @Description Downloads a file from SFTP and displays its content
// @Tags sftp
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /protected/sftp [get]


func SFTPHandler(c *gin.Context) {
	if err := services.DownloadAndProcessFile(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File downloaded and processed successfully"})
}
