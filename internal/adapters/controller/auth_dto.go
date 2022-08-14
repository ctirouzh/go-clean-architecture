package controller

import (
	"lms/internal/domain/user"
	"time"
)

type SignUpRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	Verified  bool      `json:"verified"`
	Banned    bool      `json:"banned"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (res *UserDTO) Prepare(user *user.User) {
	res.ID = user.ID.String()
	res.Username = user.Username
	res.Email = user.Email
	res.Type = user.Type.String()
	res.Verified = user.Verified
	res.Banned = user.Banned
	res.CreatedAt = user.CreatedAt
	res.UpdatedAt = user.UpdatedAt
}
