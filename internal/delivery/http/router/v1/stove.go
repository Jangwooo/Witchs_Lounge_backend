package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
)

func NewStoveRouter(app *fiber.App, stoveHandler *handler.StoveHandler) {
	stove := app.Group("/api/v1/stove")

	stove.Post("/signin", stoveHandler.SignIn)
}
