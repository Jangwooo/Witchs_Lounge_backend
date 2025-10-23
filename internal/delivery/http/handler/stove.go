package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/usecase"
)

// StoveHandler Stove 전용 핸들러
type StoveHandler struct {
	stoveUseCase usecase.StoveUseCase
}

// NewStoveHandler Stove 핸들러 생성자
func NewStoveHandler(stoveUseCase usecase.StoveUseCase) *StoveHandler {
	return &StoveHandler{
		stoveUseCase: stoveUseCase,
	}
}

// SignIn Stove 로그인 처리
func (h *StoveHandler) SignIn(c *fiber.Ctx) error {
	req := c.Locals("body").(entity.StoveSignInRequest)

	// Stove 로그인 처리
	sessionResp, err := h.stoveUseCase.SignInWithStove(c.Context(), struct {
		ID          string
		Email       string
		AvatarUrl   string
		DisplayName string
	}{
		ID:          req.ID,
		Email:       req.Email,
		AvatarUrl:   req.AvatarUrl,
		DisplayName: req.DisplayName,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sessionResp)
}
