package controllers

import (
	"net/http"
	"strings"

	"ayuphone_api/internal/models"
	"ayuphone_api/pkg/utils"

	"github.com/gin-gonic/gin"
)

func (ac ApiController) Register(c *gin.Context) {
	var requestData struct {
		InviteCode  string `json:"invite_code,omitempty"`
		FirstName   string `json:"first_name,omitempty"`
		LastName    string `json:"last_name,omitempty"`
		Email       string `json:"email,omitempty"`                 // Can be either email or phone number
		PhoneNumber string `json:"phone_number" binding:"required"` // Can be either email or phone number
		Password    string `json:"password" binding:"required"`
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

	user := models.User{
		FirstName:   requestData.FirstName,
		LastName:    requestData.LastName,
		Email:       requestData.Email,
		PhoneNumber: requestData.PhoneNumber,
		Password:    requestData.Password,
	}

	_, err = ac.DbService.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (ac ApiController) Login(c *gin.Context) {
	var requestData struct {
		Username string `json:"username" binding:"required"` // Can be either email or phone number
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var user *models.User
	var err error

	if strings.Contains(requestData.Username, "@") {
		user, err = ac.DbService.GetUserByEmail(c, requestData.Username)
	} else {
		if err = utils.ValidatePhilippinePhoneNumber(requestData.Username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err = ac.DbService.GetUserByPhoneNumber(c, requestData.Username)
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Validate the password
	if !utils.CheckPasswordHash(requestData.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
