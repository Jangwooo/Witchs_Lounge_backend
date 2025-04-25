package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
	"github.com/witchs-lounge_backend/internal/delivery/http/middleware"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
)

// @Summary Create a new user
// @Description Creates a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User information"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /users [post]
func NewUserRouter(app *fiber.App, userHandler *handler.UserHandler, sessionStore session.SessionStore) {
	users := app.Group("/api/v1/users")

	// 공개 API 엔드포인트
	users.Post("/", userHandler.SignIn)

	// 인증이 필요한 API 엔드포인트
	authUsers := users.Group("", middleware.AuthMiddleware(sessionStore))
	authUsers.Get("/me", userHandler.GetMe)
}
