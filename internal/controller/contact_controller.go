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

// AddContact godoc
// @Summary Add a contact
// @Description Adds a contact to the user's list
// @Tags Contacts
// @Accept json
// @Produce json
// @Param contact body map[string]int true "Contact ID"
// @Success 200 {object} map[string]string
// @Router /contacts/add [post]
func (cc *ContactController) AddContact(context *gin.Context) {
	var requestBody struct {
		ContactID int64 `json:"contact_id"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	userID := context.Keys["userID"].(int64)

	err := cc.service.AddContact(userID, requestBody.ContactID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Successfully added the contact"})
}

// GetContacts godoc
// @Summary Get all contacts
// @Description Retrieves all contacts for a user
// @Tags Contacts
// @Accept json
// @Produce json
// @Param user_id query int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /contacts/list [get]
func (cc *ContactController) GetContacts(context *gin.Context) {
	// var requestBody struct {
	// 	UserID int `json:"user_id"`
	// }
	// if err := context.ShouldBindQuery(&requestBody); err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
	// 	return
	// }

	// log.Printf("User id from context is: %v", context.Keys["userID"])
	userID := context.Keys["userID"].(int64)

	// log.Printf("User id is: %d", userID)
	contacts, err := cc.service.GetContacts(userID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

// DeleteContact godoc
// @Summary Delete a contact
// @Description Deletes a contact from the user's list
// @Tags Contacts
// @Accept json
// @Produce json
// @Param contact body map[string]int true "Contact ID"
// @Success 200 {object} map[string]string
// @Router /contacts/remove/:id [delete]
func (cc *ContactController) DeleteContact(context *gin.Context) {
	var requestBody struct {
		ContactID int64 `json:"contact_id"`
	}
	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request"})
		return
	}
	userID := context.Keys["userID"].(int64)

	err := cc.service.RemoveContact(userID, requestBody.ContactID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully deleted the contact"})
}
