package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type MockUserUseCase struct {
	mock.Mock
}

func (m *MockUserUseCase) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error) {
	args := m.Called(ctx, appID, ticket)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.User), args.Error(1)
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

	tests := []struct {
		name           string
		appID          string
		ticket         string
		mockUser       *entity.User
		mockError      error
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:   "성공_로그인",
			appID:  "730",
			ticket: "valid_ticket",
			mockUser: &entity.User{
				User: &ent.User{
					SteamID: "12345",
				},
			},
			mockError:      nil,
			expectedStatus: fiber.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":            "00000000-0000-0000-0000-000000000000",
				"steam_id":      "12345",
				"created_at":    "0001-01-01T00:00:00Z",
				"updated_at":    "0001-01-01T00:00:00Z",
				"last_login_at": "0001-01-01T00:00:00Z",
			},
		},
		{
			name:           "실패_잘못된_티켓",
			appID:          "730",
			ticket:         "invalid_ticket",
			mockUser:       nil,
			mockError:      fiber.NewError(fiber.StatusBadRequest, "invalid ticket"),
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"message": "Failed to create user",
				"error":   "invalid ticket",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock setup
			mockUseCase.On("VerifyAppTicket", mock.Anything, tt.appID, tt.ticket).Return(tt.mockUser, tt.mockError)

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
			assert.Equal(t, tt.expectedBody, result)
		})
	}
}
