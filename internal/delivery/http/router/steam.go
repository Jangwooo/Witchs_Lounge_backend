package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
)

func NewSteamRouter(app *fiber.App, steamHandler *handler.SteamHandler) {
	steam := app.Group("/api/v1/steam")
	users := steam.Group("/users")

	users.Get("/:steam_id", steamHandler.GetUserInfo)
	users.Get("/:steam_id/friends", steamHandler.GetUserFriends)
	users.Get("/:steam_id/games", steamHandler.GetUserGames)
	users.Get("/:steam_id/bans", steamHandler.IsUserBanned)
}
