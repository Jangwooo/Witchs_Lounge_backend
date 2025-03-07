package entity

import (
	"time"

	"github.com/witchs-lounge_backend/ent"
)

// User represents the domain entity for a user
type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest represents the request for creating a new user
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name" validate:"required"`
}

// UserResponse represents the response for user data
type UserResponse struct {
	ID        int       `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// FromEntUser converts an ent.User to our domain User
func FromEntUser(entUser *ent.User) *User {
	return &User{
		ID:        entUser.ID,
		Email:     entUser.Email,
		Password:  entUser.Password,
		Name:      entUser.Name,
		CreatedAt: entUser.CreatedAt,
		UpdatedAt: entUser.UpdatedAt,
	}
}

// ToEntUser converts our domain User to ent.User creation parameters
func (u *User) ToEntUser(client *ent.Client) *ent.UserCreate {
	return client.User.Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetName(u.Name)
}
