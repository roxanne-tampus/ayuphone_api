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
		RoleID      int    `json:"role_id" binding:"required"`
		Password    string `json:"password" binding:"required"`
	}

	organizationID := c.Param("organization_id")
	orgID, err := strconv.ParseInt(organizationID, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid organization ID")
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to hash password")
		return
	}
	requestData.Password = hashedPassword

	if requestData.PhoneNumber != "" {
		if err := utils.ValidatePhilippinePhoneNumber(requestData.PhoneNumber); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
	}

	if !ac.CheckRoleID(c, 1) && !ac.CheckRoleID(c, 2) {
		utils.ErrorResponse(c, http.StatusUnauthorized, " Unauthorized")
		return
	}

	userIdString, _ := c.Get("user_id")
	currentUser := userIdString.(int64)

	if ac.CheckRoleID(c, 2) {
		_, err := ac.DbService.GetOrganizationByUserID(c, orgID, currentUser)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, " Unauthorized")
			return
		}
	}

	user := models.User{
		FirstName:   requestData.FirstName,
		LastName:    requestData.LastName,
		Email:       requestData.Email,
		PhoneNumber: requestData.PhoneNumber,
		RoleID:      requestData.RoleID,
		Password:    requestData.Password,
	}

	userId, err := ac.DbService.CreateUser(c, &user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	organizationUser := models.OrganizationUser{
		UserID:         userId,
		OrganizationID: orgID,
		InvitedBy:      &currentUser,
		RoleID:         requestData.RoleID,
		Status:         "approved",
	}

	orgUserSuccess, err := ac.DbService.CreateOrganizationUser(c, &organizationUser)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to create OrganizationUser")
		return
	}

	utils.JSONResponse(c, true, "Organization user "+requestData.FirstName+" is created successfully", orgUserSuccess)
}

// GetOrganizationUser retrieves a OrganizationUser by ID
func (ac ApiController) GetOrganizationUsers(c *gin.Context) {
	filter := c.Query("filter")
	organizationID := c.Param("organization_id")
	orgID, err := strconv.ParseInt(organizationID, 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid organization ID")
		return
	}

	if !ac.CheckRoleID(c, 1) && !ac.CheckRoleID(c, 2) {
		utils.ErrorResponse(c, http.StatusUnauthorized, " Unauthorized")
		return
	}

	userIdString, _ := c.Get("user_id")
	currentUser := userIdString.(int64)

	if ac.CheckRoleID(c, 2) {
		_, err := ac.DbService.GetOrganizationByUserID(c, orgID, currentUser)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, " Unauthorized")
			return
		}
	}

	organizations, err := ac.DbService.GetOrganizationUsers(c, orgID, filter)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to retrieve OrganizationUser")
		return
	}

	utils.JSONResponse(c, true, "", organizations)
}

// GetOrganizationUser retrieves a OrganizationUser by ID
func (ac ApiController) GetOrganizationUser(c *gin.Context) {
	organizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid OrganizationUser ID")
		return
	}

	organizationUser, err := ac.DbService.GetOrganizationUserByID(c, organizationUserID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to retrieve OrganizationUser")
		return
	}
	utils.JSONResponse(c, true, "All organization users", organizationUser)
}

// UpdateOrganizationUser handles updating an existing OrganizationUser
func (ac ApiController) UpdateOrganizationUser(c *gin.Context) {
	organizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid OrganizationUser ID")
		return
	}

	var organizationUser models.OrganizationUser
	if err := c.ShouldBindJSON(&organizationUser); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	organizationUser.ID = organizationUserID
	if err := ac.DbService.UpdateOrganizationUser(c, &organizationUser); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to update OrganizationUser")
		return
	}
	utils.JSONResponse(c, true, "Organization user updated", organizationUser)
}

// DeleteOrganizationUser handles the deletion of a OrganizationUser by ID
func (ac ApiController) DeleteOrganizationUser(c *gin.Context) {
	organizationUserID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid OrganizationUser ID")
		return
	}

	if err := ac.DbService.DeleteOrganizationUser(c, organizationUserID); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Failed to delete OrganizationUser")
		return
	}

	utils.JSONResponse(c, true, "Organization user deleted", nil)
}
