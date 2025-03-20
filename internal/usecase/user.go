package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type UserUseCase interface {
	VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error)
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

func (u *userUseCase) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error) {
	return u.userRepo.VerifyAppTicket(ctx, appID, ticket)
}

func (u *userUseCase) FindBySteamID(ctx context.Context, steamID string) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindBySteamID(ctx, steamID)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}

func (u *userUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}
