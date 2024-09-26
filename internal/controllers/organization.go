package controllers

import (
	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrganization handles the creation of a new Organization
func (ac ApiController) CreateOrganization(c *gin.Context) {
	var requestData struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	exists := ac.CheckRole(c, "superadmin")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
		return
	}

	organization := &models.Organization{
		Name: requestData.Name,
	}

	organization, err := ac.DbService.CreateOrganization(c, organization)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to create Organization")
		return
	}

	utils.JSONResponse(c, true, "Organization "+requestData.Name+" is created successfully", organization)
}

// GetOrganization retrieves a Organization by ID
func (ac ApiController) GetOrganizations(c *gin.Context) {
	exists := ac.CheckRole(c, "superadmin")
	if !exists {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
		return
	}

	organizations, err := ac.DbService.GetOrganizations(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to retrieve Organization")
		return
	}

	utils.JSONResponse(c, true, "", organizations)
}

// GetOrganization retrieves a Organization by ID
func (ac ApiController) GetOrganization(c *gin.Context) {
	organizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid Organization ID")
		return
	}

	organization, err := ac.DbService.GetOrganizationByID(c, organizationID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to retrieve Organization")
		return
	}

	utils.JSONResponse(c, true, "", organization)
}

// UpdateOrganization handles updating an existing Organization
func (ac ApiController) UpdateOrganization(c *gin.Context) {
	organizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid Organization ID")
		return
	}

	var organization models.Organization
	if err := c.ShouldBindJSON(&organization); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	organization.ID = organizationID
	if err := ac.DbService.UpdateOrganization(c, &organization); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to update Organization")
		return
	}

	utils.JSONResponse(c, true, "Organization is updated", organization)
}

// DeleteOrganization handles the deletion of a Organization by ID
func (ac ApiController) DeleteOrganization(c *gin.Context) {
	organizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid Organization ID")
		return
	}

	if err := ac.DbService.DeleteOrganization(c, organizationID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to delete OrganizationUser")
		return
	}

	utils.JSONResponse(c, true, "Organization is deleted", nil)
}
