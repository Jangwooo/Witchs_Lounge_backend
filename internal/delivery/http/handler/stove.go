package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/usecase"
)

// StoveSignInRequest Stove 로그인 요청 구조체
type StoveSignInRequest struct {
	ID          string `validate:"required" json:"id,omitempty"`
	Email       string `validate:"required" json:"email,omitempty"`
	AvatarUrl   string `validate:"required" json:"avatar_url,omitempty"`
	DisplayName string `validate:"required" json:"display_name,omitempty"`
}

func (r StoveSignInRequest) ConvertToStoveInfo() usecase.StoveInfo {
	return usecase.StoveInfo{
		ID:          r.ID,
		Email:       r.Email,
		AvatarUrl:   r.AvatarUrl,
		DisplayName: r.DisplayName,
	}
}

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
	var req StoveSignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "잘못된 요청 형식입니다",
		})
	}

	// Stove 로그인 처리
	sessionResp, err := h.stoveUseCase.SignInWithStove(c.Context(), req.ConvertToStoveInfo())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sessionResp)
}
