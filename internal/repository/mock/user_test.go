package mock

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/witchs-lounge_backend/ent"
)

func TestUserRepository_Create(t *testing.T) {
	mockRepo := new(UserRepository)
	ctx := context.Background()

	expectedUser := &ent.User{
		ID:        uuid.New(),
		Nickname:  "test",
		SteamID:   "1234567890",
	}

	mockRepo.On("Create", ctx, "test@example.com", "password123", "Test User").
		Return(expectedUser, nil)

	user, err := mockRepo.Create(ctx, "test@example.com", "password123", "Test User")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}

func TestUserRepository_FindAll(t *testing.T) {
	mockRepo := new(UserRepository)
	ctx := context.Background()

	expectedUsers := []*ent.User{
		{
			ID:        uuid.New(),
			Nickname:  "test1",
			SteamID:   "1234567890",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Nickname:  "test2",
			SteamID:   "1234567890",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("FindAll", ctx).
		Return(expectedUsers, nil)

	users, err := mockRepo.FindAll(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
	mockRepo.AssertExpectations(t)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	mockRepo := new(UserRepository)
	ctx := context.Background()

	expectedUser := &ent.User{
		ID:        uuid.New(),
		Nickname:  "test",
		SteamID:   "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("FindByEmail", ctx, "test@example.com").
		Return(expectedUser, nil)

	user, err := mockRepo.FindByEmail(ctx, "test@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, user)
	mockRepo.AssertExpectations(t)
}
