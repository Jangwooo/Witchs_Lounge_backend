package repository

import (
	"context"

	"github.com/witchs-lounge_backend/internal/domain/entity"
)

type SteamRepository interface {
	GetUserInfo(ctx context.Context, steamID string) (*entity.SteamUser, error)
	GetUserFriends(ctx context.Context, steamID string) ([]*entity.SteamUser, error)
	GetUserGames(ctx context.Context, steamID string) ([]string, error)
	GetUserBans(ctx context.Context, steamID string) (bool, error)
}
