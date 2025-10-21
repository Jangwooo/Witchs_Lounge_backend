package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/usecase"
)

// StoveSignInRequest Stove 로그인 요청 구조체
type StoveSignInRequest struct {
	Token string `json:"token" validate:"required"`
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

	// 필수 필드 검증
	if req.Token == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "토큰이 필요합니다",
		})
	}

	// Stove 로그인 처리
	sessionResp, err := h.stoveUseCase.VerifyToken(c.Context(), req.Token)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sessionResp)
}