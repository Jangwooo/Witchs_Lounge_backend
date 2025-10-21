package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

// UserUseCase 사용자 관련 UseCase
type UserUseCase interface {
	FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error)
}

type userUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase UserUseCase 생성자
func NewUserUseCase(userRepo repository.UserRepository) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

// FindByID ID로 사용자 조회
func (u *userUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}
