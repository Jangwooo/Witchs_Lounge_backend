package usecase

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
)

// StoveUseCase Stove 인증 전용 UseCase
type StoveUseCase interface {
	VerifyToken(ctx context.Context, token string) (*entity.SessionResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error)
}

type stoveUseCase struct {
	userRepo     repository.UserRepository
	sessionStore session.SessionStore
	// TODO: Stove 인증 전략 추가 필요
}

// NewStoveUseCase Stove UseCase 생성자
func NewStoveUseCase(userRepo repository.UserRepository, sessionStore session.SessionStore) StoveUseCase {
	return &stoveUseCase{
		userRepo:     userRepo,
		sessionStore: sessionStore,
	}
}

// VerifyToken Stove 토큰 검증 및 세션 생성 (로직 비움)
func (u *stoveUseCase) VerifyToken(ctx context.Context, token string) (*entity.SessionResponse, error) {
	// TODO: Stove 인증 로직 구현 필요
	return nil, fmt.Errorf("Stove 인증이 아직 구현되지 않았습니다")
}

// FindByID ID로 사용자 조회
func (u *stoveUseCase) FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user.ToResponse(), nil
}
