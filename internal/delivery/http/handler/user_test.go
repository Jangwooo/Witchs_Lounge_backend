package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.SessionResponse, error) {
	args := m.Called(ctx, appID, ticket)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.SessionResponse), args.Error(1)
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

func setupTest() (*fiber.App, *MockUserUseCase) {
	app := fiber.New()
	mockUseCase := new(MockUserUseCase)
	handler := NewUserHandler(mockUseCase)

	// Setup routes
	app.Post("/api/v1/users/signin", handler.SignIn)

	return app, mockUseCase
}

func TestUserHandler_SignIn(t *testing.T) {
	app, mockUseCase := setupTest()

	// UUID 생성
	testUUID := uuid.New()
	testUUID2 := uuid.New()
	testTime := time.Now()

	tests := []struct {
		name           string
		appID          string
		ticket         string
		expectedStatus int
		expectedBody   map[string]interface{}
		expectCall     func(mockUseCase *MockUserUseCase)
	}{
		{
			name:           "성공_로그인",
			appID:          "123456",
			ticket:         "76561199380928730", // 티켓 값으로 steam_id 사용
			expectedStatus: fiber.StatusCreated,
			expectCall: func(mockUseCase *MockUserUseCase) {
				userResponse := &entity.UserResponse{
					ID:             testUUID,
					SteamID:        "76561199380928730",
					Nickname:       "User_76561",
					SteamAvatarURL: "",
					CreatedAt:      testTime,
					UpdatedAt:      testTime,
				}
				sessionResponse := &entity.SessionResponse{
					SessionID: "test-session-id",
					User:      *userResponse,
				}
				mockUseCase.On("VerifyAppTicket", mock.Anything, "123456", "76561199380928730").Return(sessionResponse, nil)
			},
			expectedBody: map[string]interface{}{
				"session_id": "test-session-id",
				"user": map[string]interface{}{
					"id":               testUUID.String(),
					"steam_id":         "76561199380928730",
					"nickname":         "User_76561",
					"steam_avatar_url": "",
					"created_at":       mock.Anything,
					"updated_at":       mock.Anything,
				},
			},
		},
		{
			name:           "성공_짧은_티켓_로그인",
			appID:          "123456",
			ticket:         "1234", // 5글자 미만의 짧은 티켓
			expectedStatus: fiber.StatusCreated,
			expectCall: func(mockUseCase *MockUserUseCase) {
				userResponse := &entity.UserResponse{
					ID:             testUUID2,
					SteamID:        "1234",
					Nickname:       "User_1234", // 전체 ID 사용
					SteamAvatarURL: "",
					CreatedAt:      testTime,
					UpdatedAt:      testTime,
				}
				sessionResponse := &entity.SessionResponse{
					SessionID: "test-session-id-2",
					User:      *userResponse,
				}
				mockUseCase.On("VerifyAppTicket", mock.Anything, "123456", "1234").Return(sessionResponse, nil)
			},
			expectedBody: map[string]interface{}{
				"session_id": "test-session-id-2",
				"user": map[string]interface{}{
					"id":               testUUID2.String(),
					"steam_id":         "1234",
					"nickname":         "User_1234",
					"steam_avatar_url": "",
					"created_at":       mock.Anything,
					"updated_at":       mock.Anything,
				},
			},
		},
		{
			name:           "실패_로그인",
			appID:          "123456",
			ticket:         "invalid_ticket",
			expectedStatus: fiber.StatusInternalServerError,
			expectCall: func(mockUseCase *MockUserUseCase) {
				mockUseCase.On("VerifyAppTicket", mock.Anything, "123456", "invalid_ticket").Return(nil, errors.New("사용자 조회 중 오류 발생"))
			},
			expectedBody: map[string]interface{}{
				"message": "Failed to create user",
				"error":   "사용자 조회 중 오류 발생",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock setup
			tt.expectCall(mockUseCase)

			// Create request
			reqBody := map[string]string{
				"appID":  tt.appID,
				"ticket": tt.ticket,
			}
			body, _ := json.Marshal(reqBody)
			req := httptest.NewRequest("POST", "/api/v1/users/signin", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Check status code
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			// Parse response
			var result map[string]interface{}
			err = json.NewDecoder(resp.Body).Decode(&result)
			assert.NoError(t, err)

			// Check response body
			if tt.name == "성공_로그인" || tt.name == "성공_짧은_티켓_로그인" {
				// 세션 ID 확인
				require.Equal(t, tt.expectedBody["session_id"], result["session_id"])

				// 사용자 정보 확인
				userResult := result["user"].(map[string]interface{})
				userExpected := tt.expectedBody["user"].(map[string]interface{})

				for k, v := range userExpected {
					if k == "created_at" || k == "updated_at" {
						require.Contains(t, userResult, k)
					} else {
						require.Equal(t, v, userResult[k])
					}
				}
			} else {
				// 에러 케이스
				for k, v := range tt.expectedBody {
					require.Equal(t, v, result[k])
				}
			}
		})
	}
}
