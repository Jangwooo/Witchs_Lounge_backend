package steam

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/witchs-lounge_backend/internal/domain/entity"
	"github.com/witchs-lounge_backend/internal/domain/repository"
)

type steamRepository struct {
	apiKey     string
	httpClient *http.Client
}

func NewSteamRepository(apiKey string) repository.SteamRepository {
	return &steamRepository{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (r *steamRepository) GetUserInfo(ctx context.Context, steamID string) (*entity.SteamUser, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=%s&steamids=%s", r.apiKey, steamID)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			Players []struct {
				SteamID      string `json:"steamid"`
				PersonaName  string `json:"personaname"`
				ProfileURL   string `json:"profileurl"`
				Avatar       string `json:"avatar"`
				AvatarMedium string `json:"avatarmedium"`
				AvatarFull   string `json:"avatarfull"`
				CountryCode  string `json:"loccountrycode"`
				TimeCreated  int64  `json:"timecreated"`
				LastLogOff   int64  `json:"lastlogoff"`
			} `json:"players"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	if len(result.Response.Players) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	player := result.Response.Players[0]
	return &entity.SteamUser{
		SteamID:      player.SteamID,
		PersonaName:  player.PersonaName,
		ProfileURL:   player.ProfileURL,
		Avatar:       player.Avatar,
		AvatarMedium: player.AvatarMedium,
		AvatarFull:   player.AvatarFull,
		CountryCode:  player.CountryCode,
		TimeCreated:  time.Unix(player.TimeCreated, 0),
		LastLogOff:   time.Unix(player.LastLogOff, 0),
	}, nil
}

func (r *steamRepository) GetUserFriends(ctx context.Context, steamID string) ([]*entity.SteamUser, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend", r.apiKey, steamID)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		FriendsList struct {
			Friends []struct {
				SteamID      string `json:"steamid"`
				Relationship string `json:"relationship"`
				FriendSince  int64  `json:"friend_since"`
			} `json:"friends"`
		} `json:"friendslist"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var friends []*entity.SteamUser
	for _, friend := range result.FriendsList.Friends {
		userInfo, err := r.GetUserInfo(ctx, friend.SteamID)
		if err != nil {
			continue
		}
		friends = append(friends, userInfo)
	}

	return friends, nil
}

func (r *steamRepository) GetUserGames(ctx context.Context, steamID string) ([]string, error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&include_appinfo=true&include_played_free_games=true", r.apiKey, steamID)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Response struct {
			Games []struct {
				AppID           int    `json:"appid"`
				Name            string `json:"name"`
				PlaytimeForever int    `json:"playtime_forever"`
			} `json:"games"`
		} `json:"response"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	var games []string
	for _, game := range result.Response.Games {
		games = append(games, game.Name)
	}

	return games, nil
}

func (r *steamRepository) GetUserBans(ctx context.Context, steamID string) (bool, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerBans/v1/?key=%s&steamids=%s", r.apiKey, steamID)

	resp, err := r.httpClient.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result struct {
		Players []struct {
			SteamID          string `json:"SteamId"`
			CommunityBanned  bool   `json:"CommunityBanned"`
			VACBanned        bool   `json:"VACBanned"`
			NumberOfVACBans  int    `json:"NumberOfVACBans"`
			DaysSinceLastBan int    `json:"DaysSinceLastBan"`
		} `json:"players"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	if len(result.Players) == 0 {
		return false, fmt.Errorf("user not found")
	}

	return result.Players[0].CommunityBanned || result.Players[0].VACBanned, nil
}
