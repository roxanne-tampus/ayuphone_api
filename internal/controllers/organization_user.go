package controllers

import (
	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrganizationUser handles the creation of a new OrganizationUser
func (ac ApiController) CreateOrganizationUser(c *gin.Context) {
	var requestData struct {
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Role        string `json:"role" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	organizationID := c.Param("organization_id")
	orgID, err := strconv.ParseInt(organizationID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization ID"})
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	requestData.Password = hashedPassword

	if requestData.PhoneNumber != "" {
		if err := utils.ValidatePhilippinePhoneNumber(requestData.PhoneNumber); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	if !ac.CheckRole(c, "superadmin") && !ac.CheckRole(c, "admin") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	userIdString, _ := c.Get("user_id")
	currentUser := userIdString.(int64)

	if ac.CheckRole(c, "admin") {
		_, err := ac.DbService.GetOrganizationByUserID(c, orgID, currentUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorized"})
			return
		}
	}

	user := models.User{
		FirstName:   requestData.FirstName,
		LastName:    requestData.LastName,
		Email:       requestData.Email,
		PhoneNumber: requestData.PhoneNumber,
		Role:        requestData.Role,
		Password:    requestData.Password,
	}

	userId, err := ac.DbService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	organizationUser := models.OrganizationUser{
		UserID:         userId,
		OrganizationID: orgID,
		InvitedBy:      &currentUser,
		Role:           requestData.Role,
		Status:         "approved",
	}

	orgUserSuccess, err := ac.DbService.CreateOrganizationUser(c, &organizationUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OrganizationUser"})
		return
	}

	c.JSON(http.StatusCreated, orgUserSuccess)
}

// GetOrganizationUser retrieves a OrganizationUser by ID
func (ac ApiController) GetOrganizationUsers(c *gin.Context) {
	filter := c.Query("filter")
	organizationID := c.Param("organization_id")
	orgID, err := strconv.ParseInt(organizationID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid organization ID"})
		return
	}

	if !ac.CheckRole(c, "superadmin") && !ac.CheckRole(c, "admin") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unauthorized"})
		return
	}

	userIdString, _ := c.Get("user_id")
	currentUser := userIdString.(int64)

	if ac.CheckRole(c, "admin") {
		_, err := ac.DbService.GetOrganizationByUserID(c, orgID, currentUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unauthorized"})
			return
		}
	}

	organizations, err := ac.DbService.GetOrganizationUsers(c, orgID, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve OrganizationUser"})
		return
	}

	c.JSON(http.StatusOK, organizations)
}

// GetOrganizationUser retrieves a OrganizationUser by ID
func (ac ApiController) GetOrganizationUser(c *gin.Context) {
	OrganizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OrganizationUser ID"})
		return
	}

	OrganizationUser, err := ac.DbService.GetOrganizationUserByID(c, OrganizationUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve OrganizationUser"})
		return
	}

	c.JSON(http.StatusOK, OrganizationUser)
}

// UpdateOrganizationUser handles updating an existing OrganizationUser
func (ac ApiController) UpdateOrganizationUser(c *gin.Context) {
	OrganizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OrganizationUser ID"})
		return
	}

	var OrganizationUser models.OrganizationUser
	if err := c.ShouldBindJSON(&OrganizationUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	OrganizationUser.ID = OrganizationUserID
	if err := ac.DbService.UpdateOrganizationUser(c, &OrganizationUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update OrganizationUser"})
		return
	}

	c.JSON(http.StatusOK, OrganizationUser)
}

// DeleteOrganizationUser handles the deletion of a OrganizationUser by ID
func (ac ApiController) DeleteOrganizationUser(c *gin.Context) {
	OrganizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid OrganizationUser ID"})
		return
	}

	if err := ac.DbService.DeleteOrganizationUser(c, OrganizationUserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete OrganizationUser"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OrganizationUser deleted"})
}
