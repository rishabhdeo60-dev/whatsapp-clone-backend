package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/rishabhdeo60-dev/whatsapp-clone/docs" // Swagger docs
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/config"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/controller"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/db"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/middleware"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/repository"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/routes"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/service"
)

func StartServer() {
	// Server starting logic goes here

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		log.Printf("Listening on %s", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("server stopped")
}

// @title WhatsApp Clone Backend API
// @version 1.0
// @description REST API documentation for WhatsApp-like backend + WebSocket backend built in Go
// @host localhost:8080
// @BasePath /api/v1/
func main() {

	// Load configuration
	cfg := &config.AppConfig{
		DBConfig:    &config.DBConfig{},
		RedisConfig: &config.RedisConfig{},
		KafkaConfig: &config.KafkaConfig{},
		WSConfig:    &config.WebSocketConfig{},
		JWTConfig:   &config.JWTConfig{},
	}

	cfg.LoadConfig()

	// Initialize DB connection
	DB := db.NewConnection(cfg.DBConfig)
	defer DB.Close()

	// migrate database
	DB.RunMigrations(context.Background())

	// Initialize repositories, services, and controllers
	userRepo := repository.NewUserRepository(DB)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(authService)
	contactRepo := repository.NewContactRepository(DB)
	contactService := service.NewContactService(contactRepo, userRepo)
	contactController := controller.NewContactController(contactService)
	authMiddleware := middleware.NewAuthMiddleware(cfg.JWTConfig.JwtSecret)

	router := gin.Default()
	routes.AuthRoutes(router, authController)
	routes.ContactRoutes(router, contactController, authMiddleware.RequireAuth())

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("🚀 Server running at http://localhost:8080")
}
