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

	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/config"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/controller"
	"github.com/rishabhdeo60-dev/whatsapp-clone/internal/db"
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

func main() {

	// Load configuration
	cfg := &config.AppConfig{
		DBConfig:    &config.DBConfig{},
		RedisConfig: &config.RedisConfig{},
		KafkaConfig: &config.KafkaConfig{},
		WSConfig:    &config.WebSocketConfig{},
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

	router := gin.Default()
	routes.AuthRoutes(router, authController)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
	log.Println("🚀 Server running at http://localhost:8080")
}
