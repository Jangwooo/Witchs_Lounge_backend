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
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) Create(ctx context.Context, req *entity.SignInRequest) (*entity.UserResponse, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.UserResponse), args.Error(1)
}

func (m *MockUserUseCase) FindBySteamID(ctx context.Context, steamID string) (*entity.UserResponse, error) {
	args := m.Called(ctx, steamID)
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
	mockUseCase := new(MockUserUseCase)
	app := fiber.New()

	testID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name           string
		requestBody    map[string]interface{}
		mockResponse   *entity.UserResponse
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "성공_유저_생성",
			requestBody: map[string]interface{}{
				"nickname": "테스트유저",
				"steam_id": "12345",
			},
			mockResponse: &entity.UserResponse{
				ID:        testID,
				Nickname:  "테스트유저",
				SteamID:   "12345",
				CreatedAt: now,
				UpdatedAt: now,
			},
			mockError:      nil,
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"id":         testID.String(),
				"nickname":   "테스트유저",
				"steam_id":   "12345",
				"created_at": now.Format(time.RFC3339),
				"updated_at": now.Format(time.RFC3339),
			},
		},
		{
			name: "실패_잘못된_요청",
			requestBody: map[string]interface{}{
				"nickname": "", // 필수 필드 누락
			},
			mockResponse:   nil,
			mockError:      errors.New("invalid request"),
			expectedStatus: fiber.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error": "invalid request",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock 설정
			if tc.mockResponse != nil {
				mockUseCase.On("Create", mock.Anything, &entity.SignInRequest{
					AppID: tc.requestBody["appID"].(string),
					Ticket:  tc.requestBody["ticket"].(string),
				}).Return(tc.mockResponse, tc.mockError)
			}

			// 테스트 요청 생성
			body, _ := json.Marshal(tc.requestBody)
			req := httptest.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			resp, _ := app.Test(req)

			// 응답 검증
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			if tc.expectedStatus == fiber.StatusOK {
				var response map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBody, response)
			} else {
				var errorResponse map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBody["error"], errorResponse["error"])
			}
		})
	}
}

func TestUserHandler_FindBySteamID(t *testing.T) {
	mockUseCase := new(MockUserUseCase)
	app := fiber.New()

	testID := uuid.New()
	now := time.Now()

	testCases := []struct {
		name           string
		steamID        string
		mockResponse   *entity.UserResponse
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:    "성공_스팀ID로_유저_찾기",
			steamID: "12345",
			mockResponse: &entity.UserResponse{
				ID:        testID,
				Nickname:  "테스트유저",
				SteamID:   "12345",
				CreatedAt: now,
				UpdatedAt: now,
			},
			mockError:      nil,
			expectedStatus: fiber.StatusOK,
			expectedBody: map[string]interface{}{
				"id":         testID.String(),
				"nickname":   "테스트유저",
				"steam_id":   "12345",
				"created_at": now.Format(time.RFC3339),
				"updated_at": now.Format(time.RFC3339),
			},
		},
		{
			name:           "실패_유저_없음",
			steamID:        "99999",
			mockResponse:   nil,
			mockError:      errors.New("user not found"),
			expectedStatus: fiber.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error": "user not found",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Mock 설정
			if tc.mockResponse != nil {
				mockUseCase.On("FindBySteamID", mock.Anything, tc.steamID).Return(tc.mockResponse, tc.mockError)
			}

			// 테스트 요청 생성
			req := httptest.NewRequest("GET", "/api/v1/users/steam/"+tc.steamID, nil)
			resp, _ := app.Test(req)

			// 응답 검증
			assert.Equal(t, tc.expectedStatus, resp.StatusCode)

			if tc.expectedStatus == fiber.StatusOK {
				var response map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&response)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBody, response)
			} else {
				var errorResponse map[string]interface{}
				err := json.NewDecoder(resp.Body).Decode(&errorResponse)
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedBody["error"], errorResponse["error"])
			}
		})
	}
}
