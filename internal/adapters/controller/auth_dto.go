package controller

import (
	"lms/internal/domain/user"
	"time"
)

type SignUpRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type SingInRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken string `json:"access_token"`
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

func (res *UserDTO) Prepare(usr user.User) {
	res.ID = usr.ID.String()
	res.Username = usr.Username
	res.Email = usr.Email
	res.Type = usr.Type.String()
	res.Verified = usr.Verified
	res.Banned = usr.Banned
	res.CreatedAt = usr.CreatedAt
	res.UpdatedAt = usr.UpdatedAt
}
