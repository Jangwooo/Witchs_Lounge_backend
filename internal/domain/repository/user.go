package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindBySteamID(ctx context.Context, steamID string) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}
