package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/service"
)

type ContactController struct {
	// Define necessary services here
	service service.ContactService
}

func NewContactController(service service.ContactService) *ContactController {
	return &ContactController{
		service: service,
	}
}

// Adds a new contact
func (cc *ContactController) AddContact(context *gin.Context) {
	var requestBody struct {
		UserID    int `json:"user_id"`
		ContactID int `json:"contact_id"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := cc.service.AddContact(requestBody.UserID, requestBody.ContactID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully added the contact"})
}

// Retrieves contacts for a user
func (cc *ContactController) GetContacts(context *gin.Context) {
	var requestBody struct {
		UserID int `json:"user_id"`
	}
	if err := context.ShouldBindQuery(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	contacts, err := cc.service.GetContacts(requestBody.UserID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, contacts)
}

// Deletes a contact
func (cc *ContactController) DeleteContact(context *gin.Context) {
	var requestBody struct {
		UserID    int `json:"user_id"`
		ContactID int `json:"contact_id"`
	}
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}

	err := cc.service.RemoveContact(requestBody.UserID, requestBody.ContactID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the contact"})
}
