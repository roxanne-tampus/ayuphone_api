package controllers

import (
	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateTransaction handles the creation of a new transaction
func (ac ApiController) CreateTransaction(c *gin.Context) {
	var requestData struct {
		DeviceId      int64   `json:"device_id" binding:"required"`
		DeviceIssueId *int64  `json:"device_issue_id,omitempty"`
		Note          *string `json:"note,omitempty"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "error: "+err.Error())
		return
	}

	customerId, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "error: Unauthorized")
		return
	}

	transaction := &models.Transaction{
		CustomerId:    customerId.(int64),
		DeviceId:      requestData.DeviceId,
		DeviceIssueId: requestData.DeviceIssueId,
		Note:          requestData.Note,
	}

	transaction, err := ac.DbService.CreateTransaction(c, transaction)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to create transaction")
		return
	}
	utils.JSONResponse(c, true, fmt.Sprintf("Transaction ID: %d", transaction.ID), nil)
}

// GetTransaction retrieves a transaction by ID
func (ac ApiController) GetTransactions(c *gin.Context) {
	filter := c.Query("filter")
	user_id, exists := c.Get("user_id")
	if !exists {
		utils.ErrorResponse(c, http.StatusUnauthorized, "error: Unauthorized")
		return
	}

	if ac.CheckRole(c, "customer") {
		transactions, err := ac.DbService.GetTransactions(c, user_id.(int64), "")
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to retrieve transaction")
			return
		}
		utils.JSONResponse(c, true, "", transactions)
	} else if ac.CheckRole(c, "technician") {
		utils.ErrorResponse(c, http.StatusUnauthorized, "error: Technicians are not authorized to view transactions")
	} else if ac.CheckRole(c, "admin") {
		if filter == "" {
			filter = "pending"
		}
		transactions, err := ac.DbService.GetTransactions(c, 0, filter)
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to retrieve transaction")
			return
		}
		utils.JSONResponse(c, true, "", transactions)
	} else {
		transactions, err := ac.DbService.GetTransactions(c, 0, "")
		if err != nil {
			utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to retrieve transaction")
			return
		}
		utils.JSONResponse(c, true, "", transactions)
	}
}

// GetTransaction retrieves a transaction by ID
func (ac ApiController) GetTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Invalid transaction ID")
		return
	}

	transaction, err := ac.DbService.GetTransactionByID(c, transactionID)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to retrieve transaction")

		return
	}
	utils.JSONResponse(c, true, "", transaction)
}

// UpdateTransaction handles updating an existing transaction
func (ac ApiController) UpdateTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Invalid transaction ID")
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "error: "+err.Error())
		return
	}

	transaction.ID = transactionID
	if err := ac.DbService.UpdateTransaction(c, &transaction); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to update transaction")
		return
	}
	utils.JSONResponse(c, true, "Transaction is updated", nil)
}

// DeleteTransaction handles the deletion of a transaction by ID
func (ac ApiController) DeleteTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Invalid transaction ID")
		return
	}

	if err := ac.DbService.DeleteTransaction(c, transactionID); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to delete transaction")
		return
	}

	utils.JSONResponse(c, true, "Transaction is deleted", nil)
}
