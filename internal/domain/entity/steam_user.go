package entity

import "time"

// SteamUser represents the Steam user information
type SteamUser struct {
	SteamID       string    `json:"steam_id"`
	PersonaName   string    `json:"persona_name"`
	ProfileURL    string    `json:"profile_url"`
	Avatar        string    `json:"avatar"`
	AvatarMedium  string    `json:"avatar_medium"`
	AvatarFull    string    `json:"avatar_full"`
	CountryCode   string    `json:"country_code"`
	TimeCreated   time.Time `json:"time_created"`
	LastLogOff    time.Time `json:"last_log_off"`
	GameID        string    `json:"game_id,omitempty"`
	GameExtraInfo string    `json:"game_extra_info,omitempty"`
}

// SteamUserResponse represents the response for Steam user data
type SteamUserResponse struct {
	SteamID      string    `json:"steam_id"`
	PersonaName  string    `json:"persona_name"`
	ProfileURL   string    `json:"profile_url"`
	Avatar       string    `json:"avatar"`
	AvatarMedium string    `json:"avatar_medium"`
	AvatarFull   string    `json:"avatar_full"`
	CountryCode  string    `json:"country_code"`
	TimeCreated  time.Time `json:"time_created"`
	LastLogOff   time.Time `json:"last_log_off"`
}
