package usecase

import (
	"context"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type UserUseCase interface {
	Create(ctx context.Context, req *entity.CreateUserRequest) (*entity.UserResponse, error)
	GetAll(ctx context.Context) ([]*entity.UserResponse, error)
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
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}

	createdUser, err := u.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return &entity.UserResponse{
		ID:        createdUser.ID,
		Email:     createdUser.Email,
		Name:      createdUser.Name,
		CreatedAt: createdUser.CreatedAt,
		UpdatedAt: createdUser.UpdatedAt,
	}, nil
}

func (u *userUseCase) GetAll(ctx context.Context) ([]*entity.UserResponse, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var responses []*entity.UserResponse
	for _, user := range users {
		responses = append(responses, &entity.UserResponse{
			ID:        user.ID,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return responses, nil
}
