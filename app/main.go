package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
	"github.com/witchs-lounge_backend/internal/delivery/http/router"
	"github.com/witchs-lounge_backend/internal/infrastructure/database"
	redisClient "github.com/witchs-lounge_backend/internal/infrastructure/redis"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
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
	mode := flag.String("mode", "prod", "mode")
	flag.Parse()

	log.Printf("Current mode is: %s\n", *mode)

	dbConfig, err := database.LoadConfig(mode)
	if err != nil {
		log.Fatalf("Failed to load database config: %v", err)
	}

	// Initialize database connection
	client, err := database.NewEntClient(dbConfig)
	if err != nil {
		log.Fatalf("Failed to create database client: %v", err)
	}
	defer client.Close()

	// Redis 연결 설정
	redisAddr := "redis:6379" // 스탠드얼론 Redis 주소

	log.Printf("Redis에 연결 중: %s", redisAddr)

	// Redis 클라이언트 생성
	redisClient := redisClient.NewRedisClient(redisAddr, "", 0)

	// Redis 연결 테스트
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("⚠️ Redis 연결 실패: %v", err)
		log.Printf("⚠️ 세션 기능이 작동하지 않을 수 있습니다.")
	} else {
		log.Printf("Redis 연결 성공!")
	}

	// 세션 스토어 초기화 (스탠드얼론 모드)
	sessionStore := session.NewRedisSessionStore(redisClient, 24*time.Hour)

	// Initialize repositories
	userRepo := repository.NewUserRepository(client)

	// Initialize use cases
	userUseCase := usecase.NewUserUseCase(userRepo, sessionStore)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userUseCase)

	// Create Fiber app
	app := fiber.New()

	app.Use(logger.New())

	// Swagger documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Health check
	app.Get("/api/v1/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":      "ok",
			"server_mode": *mode,
			"redis_mode":  "standalone",
		})
	})

	// Initialize routers
	router.NewUserRouter(app, userHandler, sessionStore)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
