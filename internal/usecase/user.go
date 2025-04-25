package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
)

type UserUseCase interface {
	VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.SessionResponse, error)
	FindBySteamID(ctx context.Context, steamID string) (*entity.UserResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error)
}

type userUseCase struct {
	userRepo     repository.UserRepository
	sessionStore session.SessionStore
}

func NewUserUseCase(userRepo repository.UserRepository, sessionStore session.SessionStore) UserUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		sessionStore: sessionStore,
	}
}

func (u *userUseCase) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.SessionResponse, error) {
	// 1. 티켓 검증 (개발 단계에서는 항상 성공으로 처리)
	// 실제로는 스팀 API를 호출하여 티켓 검증 및 사용자 정보를 가져오는 로직이 들어갈 예정
	// 현재는 ticket을 steam_id로 사용
	steamID := ticket

	// 2. steam_id로 사용자 조회
	existingUser, err := u.userRepo.FindBySteamID(ctx, steamID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("사용자 조회 중 오류 발생: %w", err)
	}

	var user *entity.User

	// 3. 사용자 존재 여부에 따른 처리
	if existingUser == nil {
		// 3-1. 사용자가 존재하지 않으면 새 사용자 생성
		nickname := "User_" + steamID
		if len(steamID) > 5 {
			nickname = "User_" + steamID[:5]
		}

		createReq := &entity.CreateUserRequest{
			SteamID:        steamID,
			Nickname:       nickname,
			SteamAvatarURL: "", // 기본값
		}

		user, err = u.userRepo.Create(ctx, createReq)
		if err != nil {
			return nil, fmt.Errorf("새 사용자 생성 중 오류 발생: %w", err)
		}
	} else {
		// 3-2. 사용자가 존재하면 해당 사용자 사용
		user = existingUser

		// 마지막 로그인 시간 업데이트 (선택적)
		_, err := u.userRepo.UpdateLastLogin(ctx, user.ID, time.Now())
		if err != nil {
			// 로그인 시간 업데이트 실패는 치명적인 오류가 아니므로 로그만 남기고 계속 진행
			fmt.Printf("마지막 로그인 시간 업데이트 중 오류 발생: %v\n", err)
		}
	}

	// 4. 세션 생성
	sessionID, err := u.sessionStore.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("세션 생성 중 오류 발생: %w", err)
	}

	// 5. 세션 응답 생성
	return user.ToSessionResponse(sessionID), nil
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
