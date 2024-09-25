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
		utils.ErrorResponse(c, http.StatusBadRequest, "error: "+err.Error())
		return
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(requestData.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to hash password")
		return
	}
	requestData.Password = hashedPassword

	if requestData.PhoneNumber != "" {
		if err := utils.ValidatePhilippinePhoneNumber(requestData.PhoneNumber); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "error: "+err.Error())
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
		utils.ErrorResponse(c, http.StatusConflict, "error: "+err.Error())
		return
	}

	utils.JSONResponse(c, true, "User registered successfully", nil)
}

func (ac ApiController) Login(c *gin.Context) {
	var requestData struct {
		Username string `json:"username" binding:"required"` // Can be either email or phone number
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, "error: Invalid input")
		return
	}

	var user *models.User
	var err error

	if strings.Contains(requestData.Username, "@") {
		user, err = ac.DbService.GetUserByEmail(c, requestData.Username)
	} else {
		if err = utils.ValidatePhilippinePhoneNumber(requestData.Username); err != nil {
			utils.ErrorResponse(c, http.StatusBadRequest, "error: "+err.Error())
			return
		}
		user, err = ac.DbService.GetUserByPhoneNumber(c, requestData.Username)
	}

	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "error: Invalid credentials")
		return
	}

	// Validate the password
	if !utils.CheckPasswordHash(requestData.Password, user.Password) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "error: Invalid credentials")
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "error: Failed to generate token")
		return
	}

	utils.JSONResponse(c, true, "User logged in successfully", gin.H{"token": token})
}
