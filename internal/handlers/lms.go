package handlers

import (
	"net/http"

	"github.com/NeginSal/api-hub/internal/services"
	"github.com/gin-gonic/gin"
)

// FetchLMSHandler godoc
// @Summary Fetch LMS data and store to MongoDB
// @Description Downloads data from external LMS API and saves to MongoDB
// @Tags lms
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /protected/lms [get]

func FetchLMSHandler(c *gin.Context) {
	if err := services.FetchAndStoreLMSData(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "LMS data stored successfully"})
}
