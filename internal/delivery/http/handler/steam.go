package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/usecase"
)

type SteamHandler struct {
	steamUseCase usecase.SteamUseCase
}

func NewSteamHandler(steamUseCase usecase.SteamUseCase) *SteamHandler {
	return &SteamHandler{
		steamUseCase: steamUseCase,
	}
}

// @Summary Get Steam user information
// @Description Get detailed information about a Steam user
// @Tags steam
// @Accept json
// @Produce json
// @Param steam_id path string true "Steam ID"
// @Success 200 {object} entity.SteamUserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/steam/users/{steam_id} [get]
func (h *SteamHandler) GetUserInfo(c *fiber.Ctx) error {
	steamID := c.Params("steam_id")
	if steamID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "steam_id is required",
		})
	}

	user, err := h.steamUseCase.GetUserInfo(c.Context(), steamID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}

// @Summary Get Steam user's friends
// @Description Get list of Steam user's friends
// @Tags steam
// @Accept json
// @Produce json
// @Param steam_id path string true "Steam ID"
// @Success 200 {array} entity.SteamUserResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/steam/users/{steam_id}/friends [get]
func (h *SteamHandler) GetUserFriends(c *fiber.Ctx) error {
	steamID := c.Params("steam_id")
	if steamID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "steam_id is required",
		})
	}

	friends, err := h.steamUseCase.GetUserFriends(c.Context(), steamID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(friends)
}

// @Summary Get Steam user's games
// @Description Get list of Steam user's owned games
// @Tags steam
// @Accept json
// @Produce json
// @Param steam_id path string true "Steam ID"
// @Success 200 {array} string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/steam/users/{steam_id}/games [get]
func (h *SteamHandler) GetUserGames(c *fiber.Ctx) error {
	steamID := c.Params("steam_id")
	if steamID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "steam_id is required",
		})
	}

	games, err := h.steamUseCase.GetUserGames(c.Context(), steamID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(games)
}

// @Summary Check if Steam user is banned
// @Description Check if a Steam user is banned (VAC or Community)
// @Tags steam
// @Accept json
// @Produce json
// @Param steam_id path string true "Steam ID"
// @Success 200 {object} map[string]bool
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /api/v1/steam/users/{steam_id}/bans [get]
func (h *SteamHandler) IsUserBanned(c *fiber.Ctx) error {
	steamID := c.Params("steam_id")
	if steamID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "steam_id is required",
		})
	}

	isBanned, err := h.steamUseCase.IsUserBanned(c.Context(), steamID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"is_banned": isBanned,
	})
}
