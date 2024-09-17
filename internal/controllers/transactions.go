package controllers

import (
	"ayuphone_api/internal/models"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerId, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

// GetTransaction retrieves a transaction by ID
func (ac ApiController) GetTransactions(c *gin.Context) {
	user_id, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	transactions, err := ac.DbService.GetTransactions(c, user_id.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}

// GetTransaction retrieves a transaction by ID
func (ac ApiController) GetTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	transaction, err := ac.DbService.GetTransactionByID(c, transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// UpdateTransaction handles updating an existing transaction
func (ac ApiController) UpdateTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.ID = transactionID
	if err := ac.DbService.UpdateTransaction(c, &transaction); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// DeleteTransaction handles the deletion of a transaction by ID
func (ac ApiController) DeleteTransaction(c *gin.Context) {
	transactionID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	if err := ac.DbService.DeleteTransaction(c, transactionID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}
