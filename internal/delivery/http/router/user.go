package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
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
func NewUserRouter(app *fiber.App, userHandler *handler.UserHandler) {
	users := app.Group("/api/v1/users")
	users.Post("/", userHandler.SignIn)
}
