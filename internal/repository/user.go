package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent"
	"github.com/witchs-lounge_backend/ent/user"
	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

// VerifyAppTicket은 Steam API를 통해 티켓을 검증하고 사용자를 반환하는 메소드입니다.
// 현재는 개발 단계이므로 항상 성공적으로 검증되었다고 가정합니다.
func (r *userRepository) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error) {
	// 실제로는 Steam API를 호출하여 티켓을 검증하고, 사용자 정보를 가져와야 합니다.
	// 개발 단계에서는 임의의 Steam ID로 처리합니다.
	testSteamID := ticket // 편의상 티켓 자체를 Steam ID로 사용

	// Steam ID로 사용자 조회
	entUser, err := r.client.User.Query().Where(user.SteamID(testSteamID)).Only(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, err
		}

		// 사용자 닉네임 생성 (안전하게 슬라이싱)
		nickname := "User_" + testSteamID
		if len(testSteamID) > 5 {
			nickname = "User_" + testSteamID[:5]
		}

		// 사용자가 없으면 새로 생성
		entUser, err = r.client.User.Create().
			SetNickname(nickname).
			SetSteamID(testSteamID).
			SetSteamAvatarURL("default").
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	return entity.NewUser(entUser), nil
}

func (r *userRepository) FindBySteamID(ctx context.Context, steamID string) (*entity.User, error) {
	entUser, err := r.client.User.Query().
		Where(user.SteamID(steamID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(entUser), nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	entUser, err := r.client.User.Query().
		Where(user.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(entUser), nil
}

// Create은 새로운 사용자를 생성하는 메소드입니다.
func (r *userRepository) Create(ctx context.Context, req *entity.CreateUserRequest) (*entity.User, error) {
	entUser, err := r.client.User.Create().
		SetNickname(req.Nickname).
		SetSteamID(req.SteamID).
		SetSteamAvatarURL(req.SteamAvatarURL).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(entUser), nil
}

// UpdateLastLogin은 사용자의 마지막 로그인 시간을 업데이트하는 메소드입니다.
func (r *userRepository) UpdateLastLogin(ctx context.Context, id uuid.UUID, lastLoginTime time.Time) (*entity.User, error) {
	entUser, err := r.client.User.UpdateOneID(id).
		SetLastLoginAt(lastLoginTime).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return entity.NewUser(entUser), nil
}
