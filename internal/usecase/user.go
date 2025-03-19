package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type UserUseCase interface {
	Create(ctx context.Context, req *entity.CreateUserRequest) (*entity.UserResponse, error)
	FindBySteamID(ctx context.Context, steamID string) (*entity.UserResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Create(ctx context.Context, req *entity.CreateUserRequest) (*entity.UserResponse, error) {
	user := &entity.User{
		Nickname: req.Nickname,
		SteamID:  req.SteamID,
	}

	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        createdUser.ID,
		Nickname:  createdUser.Nickname,
		SteamID:   createdUser.SteamID,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (u *userUseCase) FindBySteamID(ctx context.Context, steamID string) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindBySteamID(ctx, steamID)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		SteamID:   user.SteamID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *userUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        user.ID,
		Nickname:  user.Nickname,
		SteamID:   user.SteamID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
