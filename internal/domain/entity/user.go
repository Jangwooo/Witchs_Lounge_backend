package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/witchs-lounge_backend/ent"
)

// User represents the domain entity for a user
type User struct {
	*ent.User // ent 엔티티 임베드
}

// SignInRequest represents the request for creating a new user
type SignInRequest struct {
	AppID  string `json:"appID"`
	Ticket string `json:"ticket"`
}

// UserResponse represents the response for user data
type UserResponse struct {
	ID             uuid.UUID `json:"id"`
	Nickname       string    `json:"nickname"`
	SteamID        string    `json:"steam_id"`
	SteamAvatarURL string    `json:"steam_avatar_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// NewUser creates a new User instance
func NewUser(entUser *ent.User) *User {
	return &User{
		User: entUser,
	}
}

// ToResponse converts User to UserResponse
func (u *User) ToResponse() *UserResponse {
	return &UserResponse{
		ID:             u.ID,
		Nickname:       u.Nickname,
		SteamID:        u.SteamID,
		SteamAvatarURL: u.SteamAvatarURL,
		CreatedAt:      u.CreatedAt,
		UpdatedAt:      u.UpdatedAt,
	}
}

// FromEntUser converts an ent.User to our domain User
func FromEntUser(entUser *ent.User) *User {
	return &User{
		User: entUser,
	}
}
