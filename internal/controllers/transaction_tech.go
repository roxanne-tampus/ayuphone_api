package controllers

import (
	"ayuphone_api/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AssignTechnician assigns a technician to a transaction
func (ac ApiController) AssignTechnician(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var assignment models.TransactionTechnician
	if err := c.ShouldBindJSON(&assignment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	assignment.TransactionID = transactionID

	if err := ac.DbService.AssignTechnician(c, &assignment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to assign technician"})
		return
	}

	c.JSON(http.StatusCreated, assignment)
}

// UnassignTechnician unassigns a technician from a transaction
func (ac ApiController) UnassignTechnician(c *gin.Context) {
	techID, err := strconv.ParseInt(c.Param("techId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid technician ID"})
		return
	}

	if err := ac.DbService.UnassignTechnician(c, techID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unassign technician"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Technician unassigned"})
}

// GetTechnicianAssignments retrieves the assignment history for a transaction
func (ac ApiController) GetTechnicianAssignments(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	assignments, err := ac.DbService.GetTechnicianAssignments(c, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve assignments"})
		return
	}

	c.JSON(http.StatusOK, assignments)
}
