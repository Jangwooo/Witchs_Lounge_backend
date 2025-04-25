package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var req entity.SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	user, err := h.userUseCase.VerifyAppTicket(c.Context(), req.AppID, req.Ticket)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
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
