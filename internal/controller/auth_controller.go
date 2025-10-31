package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/model"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/service"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) *AuthController {
	return &AuthController{
		AuthService: authService,
	}
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (controller *AuthController) Register(context *gin.Context) {
	// Registration handler logic goes here
	var requestBody struct {
		Username     string `json:"username"`
		MobileNumber uint64 `json:"mobile_number"`
		Email        string `json:"email"`
		Name         string `json:"name"`
		Password     string `json:"password"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	err := controller.AuthService.Register(&model.User{
		Username:     requestBody.Username,
		MobileNumber: requestBody.MobileNumber,
		Email:        requestBody.Email,
		Name:         requestBody.Name,
		Password:     requestBody.Password,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login godoc
// @Summary Login user
// @Description Logs in a user and returns JWT
// @Tags Auth
// @Accept  json
// @Produce  json
// @Success 200 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (controller *AuthController) Login(context *gin.Context) {
	// Login handler logic goes here
	var requestBody struct {
		MobileUsernameEmail string `json:"mobile_username_email"`
		Password            string `json:"password"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var token string
	token, err := controller.AuthService.Login(requestBody.MobileUsernameEmail, requestBody.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	if token != "" {
		context.JSON(http.StatusOK, gin.H{"JWT token": token})
		return
	}
}
