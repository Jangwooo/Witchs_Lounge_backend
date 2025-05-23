package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
	"github.com/witchs-lounge_backend/internal/domain/strategy"
	"github.com/witchs-lounge_backend/internal/infrastructure/session"
)

type UserUseCase interface {
	VerifyAppTicket(ctx context.Context, platformType, appID, ticket string) (*entity.SessionResponse, error)
	FindByPlatformUserID(ctx context.Context, platformType, platformUserID string) (*entity.UserResponse, error)
	FindByID(ctx context.Context, id uuid.UUID) (*entity.UserResponse, error)
}

type userUseCase struct {
	userRepo     repository.UserRepository
	sessionStore session.SessionStore
	authFactory  strategy.PlatformAuthFactory
}

func NewUserUseCase(userRepo repository.UserRepository, sessionStore session.SessionStore, authFactory strategy.PlatformAuthFactory) UserUseCase {
	return &userUseCase{
		userRepo:     userRepo,
		sessionStore: sessionStore,
		authFactory:  authFactory,
	}
}

func (u *userUseCase) VerifyAppTicket(ctx context.Context, platformType, appID, ticket string) (*entity.SessionResponse, error) {
	// 1. 플랫폼 전략 가져오기
	authStrategy, err := u.authFactory.GetStrategy(platformType)
	if err != nil {
		return nil, fmt.Errorf("플랫폼 전략을 가져오는 중 오류 발생: %w", err)
	}

	// 2. 티켓 검증 및 플랫폼 사용자 정보 가져오기
	platformUserInfo, err := authStrategy.VerifyTicket(ctx, appID, ticket)
	if err != nil {
		return nil, fmt.Errorf("티켓 검증 중 오류 발생: %w", err)
	}

	// 3. 플랫폼 사용자 ID로 기존 사용자 조회
	existingUser, err := u.userRepo.FindByPlatformUserID(ctx, platformUserInfo.PlatformType, platformUserInfo.PlatformUserID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, fmt.Errorf("사용자 조회 중 오류 발생: %w", err)
	}

	var user *entity.User

	// 4. 사용자 존재 여부에 따른 처리
	if existingUser == nil {
		// 4-1. 사용자가 존재하지 않으면 새 사용자 생성
		nickname := "User_" + platformUserInfo.PlatformUserID
		if len(platformUserInfo.PlatformUserID) > 5 {
			nickname = "User_" + platformUserInfo.PlatformUserID[:5]
		}

		createReq := &entity.CreateUserRequest{
			PlatformType:        platformUserInfo.PlatformType,
			PlatformUserID:      platformUserInfo.PlatformUserID,
			PlatformEmail:       platformUserInfo.Email,
			PlatformAvatarURL:   platformUserInfo.AvatarURL,
			PlatformDisplayName: platformUserInfo.DisplayName,
			Language:            platformUserInfo.Language,
			PlatformData:        platformUserInfo.PlatformData,
			IsVerified:          platformUserInfo.IsVerified,
			Nickname:            nickname,
		}

		user, err = u.userRepo.Create(ctx, createReq)
		if err != nil {
			return nil, fmt.Errorf("새 사용자 생성 중 오류 발생: %w", err)
		}
	} else {
		// 4-2. 사용자가 존재하면 해당 사용자 사용
		user = existingUser

		// 마지막 로그인 시간 업데이트
		_, err := u.userRepo.UpdateLastLogin(ctx, user.ID, time.Now())
		if err != nil {
			// 로그인 시간 업데이트 실패는 치명적인 오류가 아니므로 로그만 남기고 계속 진행
			fmt.Printf("마지막 로그인 시간 업데이트 중 오류 발생: %v\n", err)
		}
	}

	// 5. 세션 생성
	sessionID, err := u.sessionStore.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("세션 생성 중 오류 발생: %w", err)
	}

	// 6. 세션 응답 생성
	return user.ToSessionResponse(sessionID), nil
}

func (u *userUseCase) FindByPlatformUserID(ctx context.Context, platformType, platformUserID string) (*entity.UserResponse, error) {
	user, err := u.userRepo.FindByPlatformUserID(ctx, platformType, platformUserID)
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
