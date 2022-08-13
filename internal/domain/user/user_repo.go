package user

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrUserAlreadyExist = errors.New("user already exist")
)

type Repository interface {
	Get(id uuid.UUID) (*User, error)
	Create(username, email, password string, userType UserType) error
	Delete(id uuid.UUID) error
}
