package controllers

import (
	"ayuphone_api/internal/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := ac.CheckRole(c, "superadmin")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organization := &models.Organization{
		Name: requestData.Name,
	}

	organization, err = ac.DbService.CreateOrganization(c, organization)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Organization"})
		return
	}

	c.JSON(http.StatusCreated, organization)
}

// GetOrganization retrieves a Organization by ID
func (ac ApiController) GetOrganizations(c *gin.Context) {
	exists, err := ac.CheckRole(c, "superadmin")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	organizations, err := ac.DbService.GetOrganizations(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Organization"})
		return
	}

	c.JSON(http.StatusOK, organizations)
}

// GetOrganization retrieves a Organization by ID
func (ac ApiController) GetOrganization(c *gin.Context) {
	OrganizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Organization ID"})
		return
	}

	Organization, err := ac.DbService.GetOrganizationByID(c, OrganizationID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve Organization"})
		return
	}

	c.JSON(http.StatusOK, Organization)
}

// UpdateOrganization handles updating an existing Organization
func (ac ApiController) UpdateOrganization(c *gin.Context) {
	OrganizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Organization ID"})
		return
	}

	var Organization models.Organization
	if err := c.ShouldBindJSON(&Organization); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Organization.ID = OrganizationID
	if err := ac.DbService.UpdateOrganization(c, &Organization); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Organization"})
		return
	}

	c.JSON(http.StatusOK, Organization)
}

// DeleteOrganization handles the deletion of a Organization by ID
func (ac ApiController) DeleteOrganization(c *gin.Context) {
	OrganizationID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Organization ID"})
		return
	}

	if err := ac.DbService.DeleteOrganization(c, OrganizationID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Organization"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Organization deleted"})
}
