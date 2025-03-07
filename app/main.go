package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
	"github.com/witchs-lounge_backend/internal/delivery/http/router"
	"github.com/witchs-lounge_backend/internal/infrastructure/database"
	"github.com/witchs-lounge_backend/internal/repository"
	"github.com/witchs-lounge_backend/internal/usecase"
)

// @title           Witch's Lounge API
// @version         1.0
// @description     Witch's Lounge Backend API Server
// @host           localhost:8080
// @BasePath       /api/v1
// @schemes        http
// @contact.name   API Support
// @contact.email  support@witchslounge.com

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load database configuration
	dbConfig, err := database.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Initialize database connection
	client, err := database.NewEntClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}
	defer client.Close()

	// Initialize repositories
	userRepo := repository.NewUserRepository(client)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Create Fiber app
	app := fiber.New()

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	// Initialize routers
	router.NewUserRouter(app, userHandler)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
