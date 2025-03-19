package entity

import (
	"time"

	"github.com/witchs-lounge_backend/ent"
)

// User represents the domain entity for a user
type User struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	SteamID   string    `json:"steam_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request for creating a new user
type CreateUserRequest struct {
	Nickname string `json:"nickname" validate:"required"`
	SteamID  string `json:"steam_id" validate:"required"`
}

// UserResponse represents the response for user data
type UserResponse struct {
	ID        string    `json:"id"`
	Nickname  string    `json:"nickname"`
	SteamID   string    `json:"steam_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromEntUser converts an ent.User to our domain User
func FromEntUser(entUser *ent.User) *User {
	return &User{
		ID:        entUser.ID.String(),
		Nickname:  entUser.Nickname,
		SteamID:   entUser.SteamID,
		CreatedAt: entUser.CreatedAt,
		UpdatedAt: entUser.UpdatedAt,
	}
}

// ToEntUser converts our domain User to ent.User creation parameters
func (u *User) ToEntUser(client *ent.Client) *ent.UserCreate {
	return client.User.Create().
		SetNickname(u.Nickname).
		SetSteamID(u.SteamID)
}
