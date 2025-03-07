package usecase

import (
	"context"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type SteamUseCase interface {
	GetUserInfo(ctx context.Context, steamID string) (*entity.SteamUserResponse, error)
	GetUserFriends(ctx context.Context, steamID string) ([]*entity.SteamUserResponse, error)
	GetUserGames(ctx context.Context, steamID string) ([]string, error)
	IsUserBanned(ctx context.Context, steamID string) (bool, error)
}

type steamUseCase struct {
	steamRepo repository.SteamRepository
}

func NewSteamUseCase(steamRepo repository.SteamRepository) SteamUseCase {
	return &steamUseCase{
		steamRepo: steamRepo,
	}
}

func (u *steamUseCase) GetUserInfo(ctx context.Context, steamID string) (*entity.SteamUserResponse, error) {
	user, err := u.steamRepo.GetUserInfo(ctx, steamID)
	if err != nil {
		return nil, err
	}

	return &entity.SteamUserResponse{
		SteamID:      user.SteamID,
		PersonaName:  user.PersonaName,
		ProfileURL:   user.ProfileURL,
		Avatar:       user.Avatar,
		AvatarMedium: user.AvatarMedium,
		AvatarFull:   user.AvatarFull,
		CountryCode:  user.CountryCode,
		TimeCreated:  user.TimeCreated,
		LastLogOff:   user.LastLogOff,
	}, nil
}

func (u *steamUseCase) GetUserFriends(ctx context.Context, steamID string) ([]*entity.SteamUserResponse, error) {
	friends, err := u.steamRepo.GetUserFriends(ctx, steamID)
	if err != nil {
		return nil, err
	}

	var response []*entity.SteamUserResponse
	for _, friend := range friends {
		response = append(response, &entity.SteamUserResponse{
			SteamID:      friend.SteamID,
			PersonaName:  friend.PersonaName,
			ProfileURL:   friend.ProfileURL,
			Avatar:       friend.Avatar,
			AvatarMedium: friend.AvatarMedium,
			AvatarFull:   friend.AvatarFull,
			CountryCode:  friend.CountryCode,
			TimeCreated:  friend.TimeCreated,
			LastLogOff:   friend.LastLogOff,
		})
	}

	return response, nil
}

func (u *steamUseCase) GetUserGames(ctx context.Context, steamID string) ([]string, error) {
	return u.steamRepo.GetUserGames(ctx, steamID)
}

func (u *steamUseCase) IsUserBanned(ctx context.Context, steamID string) (bool, error) {
	return u.steamRepo.GetUserBans(ctx, steamID)
}
