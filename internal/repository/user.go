package repository

import (
	"context"

	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/ent/user"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	FindAll(ctx context.Context) ([]*entity.User, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
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

func (r *userRepository) FindAll(ctx context.Context) ([]*entity.User, error) {
	entUsers, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	var users []*entity.User
	for _, entUser := range entUsers {
		users = append(users, entity.FromEntUser(entUser))
	}

	return users, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	entUser, err := r.client.User.Query().
		Where(user.Email(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.FromEntUser(entUser), nil
}
