package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
	"github.com/witchs-lounge_backend/internal/delivery/http/middleware"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

func NewStoveRouter(app *fiber.App, stoveHandler *handler.StoveHandler) {
	stove := app.Group("/api/v1/stove")

	stove.Post("/signin", middleware.ValidateBody[entity.StoveSignInRequest](), stoveHandler.SignIn)
}
