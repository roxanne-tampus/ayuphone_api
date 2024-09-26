package controllers

import (
	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AssignTechnician assigns a technician to a transaction
func (ac ApiController) AssignTechnician(c *gin.Context) {
	var requestData struct {
		TechnicianID int64 `json:"technician_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid transaction ID")
		return
	}

	if !ac.CheckRole(c, "superadmin") && !ac.CheckRole(c, "admin") {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
		return
	}

	userIdString, _ := c.Get("user_id")
	currentUser := userIdString.(int64)

	if ac.CheckRole(c, "admin") {
		organizationUser, err := ac.DbService.GetOrganizationByUserID(c, 0, currentUser)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
			return
		}

		organizationUserTech, err := ac.DbService.GetOrganizationByUserID(c, 0, requestData.TechnicianID)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
			return
		}

		if organizationUserTech.OrganizationID != organizationUser.OrganizationID {
			utils.ErrorResponse(c, http.StatusInternalServerError, " Unauthorized")
			return
		}
	}

	var assignment = models.TransactionTechnician{
		TransactionID: transactionID,
		TechnicianID:  requestData.TechnicianID,
	}

	if err := ac.DbService.AssignTechnician(c, &assignment); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Failed to assign technician")
		return
	}

	utils.JSONResponse(c, true, "Technician assigned", assignment)
}

// UnassignTechnician unassigns a technician from a transaction
func (ac ApiController) UnassignTechnician(c *gin.Context) {
	techID, err := strconv.ParseInt(c.Param("techId"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Invalid technician ID")
		return
	}

	if err := ac.DbService.UnassignTechnician(c, techID); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Failed to assign technician")
		return
	}

	utils.JSONResponse(c, true, "Technician unassigned", nil)
}

// GetTechnicianAssignments retrieves the assignment history for a transaction
func (ac ApiController) GetTechnicianAssignments(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, " Invalid transaction ID")
		return
	}

	assignments, err := ac.DbService.GetTechnicianAssignments(c, transactionID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, " Failed to assign technician")
		return
	}

	utils.JSONResponse(c, true, "Transaction assignment history", assignments)
}
