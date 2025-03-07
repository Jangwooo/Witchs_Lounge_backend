package repository

import (
	"context"

	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
