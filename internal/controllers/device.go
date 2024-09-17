package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac ApiController) GetAllDevice(c *gin.Context) {
	device, err := ac.DbService.GetAllDevice(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve device"})
		return
	}

	c.JSON(http.StatusOK, device)
}

func (ac ApiController) GetAllDeviceIssues(c *gin.Context) {
	device, err := ac.DbService.GetAllDeviceIssues(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve device"})
		return
	}

	c.JSON(http.StatusOK, device)
}
