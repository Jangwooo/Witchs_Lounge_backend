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

// CreateUserRequest 새 사용자 생성 요청 데이터
type CreateUserRequest struct {
	SteamID        string `json:"steam_id"`
	Nickname       string `json:"nickname"`
	SteamAvatarURL string `json:"steam_avatar_url"`
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

// SessionResponse 세션 정보와 유저 정보를 담는 응답 구조체
type SessionResponse struct {
	SessionID string       `json:"session_id"`
	User      UserResponse `json:"user"`
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

// ToSessionResponse converts User and sessionID to SessionResponse
func (u *User) ToSessionResponse(sessionID string) *SessionResponse {
	return &SessionResponse{
		SessionID: sessionID,
		User:      *u.ToResponse(),
	}
}

// FromEntUser converts an ent.User to our domain User
func FromEntUser(entUser *ent.User) *User {
	return &User{
		User: entUser,
	}
}
