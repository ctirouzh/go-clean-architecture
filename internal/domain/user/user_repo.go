package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrEmailAlreadyTaken    = errors.New("email already taken")
)

type Repository interface {
	Get(id uuid.UUID) (*User, error)
	Create(username, email, password string, userType UserType) (*User, error)
	Delete(id uuid.UUID) error
}
