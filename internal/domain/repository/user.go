package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type UserRepository interface {
	FindBySteamID(ctx context.Context, steamID string) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error)
}
