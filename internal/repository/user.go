package repository

import (
	"context"

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

func (r *userRepository) VerifyAppTicket(ctx context.Context, appID, ticket string) (*entity.User, error) {
	// TODO: Steam API 연동 후 유저 검증 로직 추가
	testSteamID := "76561199380928730"

	entUser, err := r.client.User.Query().Where(user.SteamID(testSteamID)).Only(ctx)
	if ent.IsNotFound(err) {
		entUser, err = r.client.User.Create().
			SetNickname("test").
			SetSteamID(testSteamID).
			SetSteamAvatarURL("default").
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, err
	}

	u := entity.FromEntUser(entUser)

	return u, nil
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
