package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/controller"
)

func AuthRoutes(router *gin.Engine, authController *controller.AuthController) {
	// Route registration logic goes here
	auth_api := router.Group("/api/v1/auth")
	{
		auth_api.POST("/register", authController.Register)
		auth_api.POST("/login", authController.Login)
	}
}

// func ProfileRoutes(router *gin.Engine, profileController *controller.ProfileController) {
// 	// Route registration logic goes here
// 	profile_api := router.Group("/api/v1/profile")
// 	{
// 		profile_api.GET("/:id", profileController.GetProfile)
// 		profile_api.PUT("/:id", profileController.UpdateProfile)
// 		profile_api.DELETE("/:id", profileController.DeleteProfile)
// 	}
// }

func ContactRoutes(router *gin.Engine, contactController *controller.ContactController, authMiddleware gin.HandlerFunc) {
	// Route registration logic goes here
	contacts_api := router.Group("/api/v1/contacts")
	contacts_api.Use(authMiddleware)
	{
		contacts_api.POST("/add", contactController.AddContact)
		contacts_api.GET("/list", contactController.GetContacts)
		contacts_api.DELETE("/remove/:id", contactController.DeleteContact)
	}
}
