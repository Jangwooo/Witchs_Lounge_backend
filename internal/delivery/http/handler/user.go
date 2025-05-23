package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/usecase"
)

// PlatformSignInRequest 플랫폼별 로그인 요청 구조체
type PlatformSignInRequest struct {
	PlatformType string `json:"platform_type" validate:"required,oneof=steam"`
	AppID        string `json:"app_id" validate:"required"`
	Ticket       string `json:"ticket" validate:"required"`
}

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

// SignIn 플랫폼별 로그인 처리
func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var req PlatformSignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "잘못된 요청 형식입니다",
		})
	}

	// 플랫폼 타입 검증
	if req.PlatformType == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "플랫폼 타입이 필요합니다",
		})
	}

	// 필수 필드 검증
	if req.AppID == "" || req.Ticket == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "앱 ID와 티켓이 필요합니다",
		})
	}

	// 플랫폼별 로그인 처리
	sessionResp, err := h.userUseCase.VerifyAppTicket(c.Context(), req.PlatformType, req.AppID, req.Ticket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(sessionResp)
}

// GetMe는 현재 인증된 사용자의 정보를 반환합니다.
func (h *UserHandler) GetMe(c *fiber.Ctx) error {
	// 미들웨어에서 설정한 사용자 정보 가져오기
	user := c.Locals("user")
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "인증 정보를 찾을 수 없습니다",
		})
	}

	// entity.User 타입으로 변환
	userData, ok := user.(*entity.User)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "사용자 정보 형식이 올바르지 않습니다",
		})
	}

	// 사용자 정보 응답
	return c.Status(fiber.StatusOK).JSON(userData.ToResponse())
}

// GetUserByPlatformID 플랫폼 사용자 ID로 사용자 조회
func (h *UserHandler) GetUserByPlatformID(c *fiber.Ctx) error {
	platformType := c.Params("platform_type")
	platformUserID := c.Params("platform_user_id")

	if platformType == "" || platformUserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "플랫폼 타입과 플랫폼 사용자 ID가 필요합니다",
		})
	}

	user, err := h.userUseCase.FindByPlatformUserID(c.Context(), platformType, platformUserID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "사용자를 찾을 수 없습니다",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
