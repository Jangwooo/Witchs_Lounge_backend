package mock

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/witchs-lounge_backend/ent"
)

func TestUserRepository_Create(t *testing.T) {
	mockRepo := new(UserRepository)
	ctx := context.Background()

	expectedUser := &ent.User{
		ID:        1,
		Email:     "test@example.com",
		Password:  "password123",
		Name:      "Test User",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
			ID:        1,
			Email:     "test1@example.com",
			Password:  "password123",
			Name:      "Test User 1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Email:     "test2@example.com",
			Password:  "password123",
			Name:      "Test User 2",
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
		ID:        1,
		Email:     "test@example.com",
		Password:  "password123",
		Name:      "Test User",
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
