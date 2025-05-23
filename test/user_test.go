package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/witchs-lounge_backend/internal/delivery/http/handler"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) VerifyAppTicket(ctx context.Context, platformType, appID, ticket string) (*entity.SessionResponse, error) {
	args := m.Called(ctx, platformType, appID, ticket)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.SessionResponse), args.Error(1)
}

func (m *MockUserUseCase) FindByPlatformUserID(ctx context.Context, platformType, platformUserID string) (*entity.UserResponse, error) {
	args := m.Called(ctx, platformType, platformUserID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserResponse), args.Error(1)
}

func (m *MockUserUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserResponse), args.Error(1)
}

func TestUserHandler_SignIn(t *testing.T) {
	testID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		mockResponse   *entity.SessionResponse
		mockError      error
		expectedStatus int
		shouldCallMock bool
	}{
		{
			name: "성공_Steam_로그인",
			requestBody: map[string]interface{}{
				"platform_type": "steam",
				"app_id":        "12345",
				"ticket":        "steam_ticket_123",
			},
			mockResponse: &entity.SessionResponse{
				SessionID: "session_123",
				User: entity.UserResponse{
					ID:                  testID,
					PlatformType:        "steam",
					PlatformUserID:      "steam_ticket_123",
					PlatformEmail:       "",
					PlatformAvatarURL:   "",
					PlatformDisplayName: "",
					Language:            "ko",
					IsVerified:          true,
					Nickname:            "User_steam",
					DisplayName:         "",
					Level:               1,
					Exp:                 0,
					Coin:                0,
					Gem:                 0,
					CreatedAt:           now,
					UpdatedAt:           now,
				},
			},
			mockError:      nil,
			expectedStatus: fiber.StatusOK,
			shouldCallMock: true,
		},
		{
			name: "실패_플랫폼_타입_누락",
			requestBody: map[string]interface{}{
				"app_id": "12345",
				"ticket": "steam_ticket_123",
			},
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: fiber.StatusBadRequest,
			shouldCallMock: false,
		},
		{
			name: "실패_앱_ID_누락",
			requestBody: map[string]interface{}{
				"platform_type": "steam",
				"ticket":        "steam_ticket_123",
			},
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: fiber.StatusBadRequest,
			shouldCallMock: false,
		},
		{
			name: "실패_티켓_누락",
			requestBody: map[string]interface{}{
				"platform_type": "steam",
				"app_id":        "12345",
			},
			mockResponse:   nil,
			mockError:      nil,
			expectedStatus: fiber.StatusBadRequest,
			shouldCallMock: false,
		},
		{
			name: "실패_UseCase_에러",
			requestBody: map[string]interface{}{
				"platform_type": "steam",
				"app_id":        "12345",
				"ticket":        "invalid_ticket",
			},
			mockResponse:   nil,
			mockError:      errors.New("티켓 검증 실패"),
			expectedStatus: fiber.StatusInternalServerError,
			shouldCallMock: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 각 테스트마다 새로운 mock과 app 생성
			mockUseCase := new(MockUserUseCase)
			userHandler := handler.NewUserHandler(mockUseCase)
			app := fiber.New()
			app.Post("/auth/signin", userHandler.SignIn)

			// Mock 설정
			if tc.shouldCallMock {
				platformType, _ := tc.requestBody["platform_type"].(string)
				appID, _ := tc.requestBody["app_id"].(string)
				ticket, _ := tc.requestBody["ticket"].(string)

				mockUseCase.On("VerifyAppTicket", mock.Anything, platformType, appID, ticket).Return(tc.mockResponse, tc.mockError).Once()
			}

			// 테스트 요청 생성
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest("POST", "/auth/signin", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			// 응답 검증
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			// Mock 호출 검증
			if tc.shouldCallMock {
				mockUseCase.AssertExpectations(t)
			}

			// 응답 내용이 있는 경우 추가 검증
			if tc.expectedStatus == fiber.StatusOK && tc.mockResponse != nil {
				var response entity.SessionResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.mockResponse.SessionID, response.SessionID)
				assert.Equal(t, tc.mockResponse.User.ID, response.User.ID)
			}
		})
	}
}

func TestUserHandler_GetUserByPlatformID(t *testing.T) {
	testID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name           string
		platformType   string
		platformUserID string
		mockResponse   *entity.UserResponse
		mockError      error
		expectedStatus int
		shouldCallMock bool
	}{
		{
			name:           "성공_플랫폼_사용자_조회",
			platformType:   "steam",
			platformUserID: "steam_123",
			mockResponse: &entity.UserResponse{
				ID:                  testID,
				PlatformType:        "steam",
				PlatformUserID:      "steam_123",
				PlatformEmail:       "",
				PlatformAvatarURL:   "",
				PlatformDisplayName: "",
				Language:            "ko",
				IsVerified:          true,
				Nickname:            "TestUser",
				DisplayName:         "",
				Level:               1,
				Exp:                 0,
				Coin:                0,
				Gem:                 0,
				CreatedAt:           now,
				UpdatedAt:           now,
			},
			mockError:      nil,
			expectedStatus: fiber.StatusOK,
			shouldCallMock: true,
		},
		{
			name:           "실패_사용자_없음",
			platformType:   "steam",
			platformUserID: "nonexistent_user",
			mockResponse:   nil,
			mockError:      errors.New("user not found"),
			expectedStatus: fiber.StatusNotFound,
			shouldCallMock: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 각 테스트마다 새로운 mock과 app 생성
			mockUseCase := new(MockUserUseCase)
			userHandler := handler.NewUserHandler(mockUseCase)
			app := fiber.New()
			app.Get("/users/:platform_type/:platform_user_id", userHandler.GetUserByPlatformID)

			// Mock 설정
			if tc.shouldCallMock {
				mockUseCase.On("FindByPlatformUserID", mock.Anything, tc.platformType, tc.platformUserID).Return(tc.mockResponse, tc.mockError).Once()
			}

			// 테스트 요청 생성
			url := "/users/" + tc.platformType + "/" + tc.platformUserID
			req := httptest.NewRequest("GET", url, nil)
			resp, _ := app.Test(req)

			// 응답 검증
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			// Mock 호출 검증
			if tc.shouldCallMock {
				mockUseCase.AssertExpectations(t)
			}

			// 성공적인 응답 내용 검증
			if tc.expectedStatus == fiber.StatusOK && tc.mockResponse != nil {
				var response entity.UserResponse
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.mockResponse.ID, response.ID)
				assert.Equal(t, tc.mockResponse.PlatformType, response.PlatformType)
				assert.Equal(t, tc.mockResponse.PlatformUserID, response.PlatformUserID)
			}
		})
	}
}
