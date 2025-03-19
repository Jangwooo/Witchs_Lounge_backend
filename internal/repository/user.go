package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/ent/user"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindBySteamID(ctx context.Context, steamID string) (*entity.User, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	entUser, err := user.ToEntUser(r.client).Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.FromEntUser(entUser), nil
}

func (r *userRepository) FindBySteamID(ctx context.Context, steamID string) (*entity.User, error) {
	entUser, err := r.client.User.Query().
		Where(user.SteamID(steamID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.FromEntUser(entUser), nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	entUser, err := r.client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.FromEntUser(entUser), nil
}
