package handler

import (
	"challenge_10/entity"
	"challenge_10/helper"
	"challenge_10/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{service}
}

func (uh *userHandler) CreateClientUser(c *gin.Context) {
	var payloadUser entity.PayloadUser

	err := c.ShouldBindJSON(&payloadUser)
	if err != nil {
		error := helper.FormatErrors(err)
		errorMessage := gin.H{"errors": error}

		response := helper.APIResponse("Unprocessable Entity!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	errorMessage := []string{}
	if payloadUser.Username == "" {
		errorMessage = append(errorMessage, "Username must be Fill!")
	}

	if payloadUser.Password == "" {
		errorMessage = append(errorMessage, "Password must be Fill!")
	}

	if payloadUser.Username == "" || payloadUser.Password == "" {
		response := helper.APIResponse("Failed to Create a User", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newClientUser, err := uh.service.CreateClientUser(payloadUser)
	if err != nil {
		response := helper.APIResponse("Failed to Create a User", http.StatusBadRequest, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes Create a User", http.StatusCreated, "Success", newClientUser)
	c.JSON(http.StatusCreated, response)
}

func (uh *userHandler) CreateAdminUser(c *gin.Context) {
	var payloadUser entity.PayloadUser

	err := c.ShouldBindJSON(&payloadUser)
	if err != nil {
		error := helper.FormatErrors(err)
		errorMessage := gin.H{"errors": error}

		response := helper.APIResponse("Unprocessable Entity!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	errorMessage := []string{}
	if payloadUser.Username == "" {
		errorMessage = append(errorMessage, "Username must be Fill!")
	}

	if payloadUser.Password == "" {
		errorMessage = append(errorMessage, "Password must be Fill!")
	}

	if payloadUser.Username == "" || payloadUser.Password == "" {
		response := helper.APIResponse("Failed to Create Admin", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newAdminUser, err := uh.service.CreateAdminUser(payloadUser)
	if err != nil {
		response := helper.APIResponse("Failed to Create Admin", http.StatusBadRequest, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes Create Admin", http.StatusCreated, "Success", newAdminUser)
	c.JSON(http.StatusCreated, response)
}

func (uh *userHandler) ValidateUser(c *gin.Context) {
	var payloadUser entity.PayloadUser

	err := c.ShouldBindJSON(&payloadUser)
	if err != nil {
		error := helper.FormatErrors(err)
		errorMessage := gin.H{"errors": error}

		response := helper.APIResponse("Unprocessable Entity!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	errorMessage := []string{}
	if payloadUser.Username == "" {
		errorMessage = append(errorMessage, "Username must be Fill!")
	}

	if payloadUser.Password == "" {
		errorMessage = append(errorMessage, "Password must be Fill!")
	}

	if payloadUser.Username == "" || payloadUser.Password == "" {
		response := helper.APIResponse("Unprocessable Entity!", http.StatusUnprocessableEntity, "Error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	getUserByUsername, err := uh.service.ValidateUser(payloadUser.Username, payloadUser.Password)
	if err != nil {
		errors := fmt.Sprintf("%s", err)
		response := helper.APIResponse("Failed to Validate User!", http.StatusUnauthorized, "Error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := helper.GenerateToken(getUserByUsername.ID, getUserByUsername.Username, getUserByUsername.Access_level)
	if err != nil {
		errors := fmt.Sprintf("%s", err)
		response := helper.APIResponse("Failed Generate Token!", http.StatusBadRequest, "Error", errors)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Validate User Success", http.StatusOK, "Success", token)
	c.JSON(http.StatusCreated, response)
}
