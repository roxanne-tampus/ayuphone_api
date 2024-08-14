package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *ApiController) LoginForTesting(c *gin.Context) {
	// Extract credentials from the request
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(c.Request.Body).Decode(&creds); err != nil {
		// Handle error
		http.Error(c.Writer, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform the login
	token, err := lc.KeyCloakClient.Login(creds.Username, creds.Password)
	if err != nil {
		log.Println(err)
		http.Error(c.Writer, "Login failed", http.StatusUnauthorized)
		return
	}

	// Send the token back in the response
	err = json.NewEncoder(c.Writer).Encode(map[string]string{"token": token})
	if err != nil {
		return
	}
}
