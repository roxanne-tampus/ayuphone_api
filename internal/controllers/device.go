package controllers

import (
	"ayuphone_api/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ac ApiController) GetAllDevice(c *gin.Context) {
	device, err := ac.DbService.GetAllDevice(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Failed to retrieve device")
		return
	}

	utils.JSONResponse(c, true, "Get all device", device)
}

func (ac ApiController) GetAllDeviceIssues(c *gin.Context) {
	device, err := ac.DbService.GetAllDeviceIssues(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Failed to retrieve device")
		return
	}
	utils.JSONResponse(c, true, "Get all device issues", device)
}
