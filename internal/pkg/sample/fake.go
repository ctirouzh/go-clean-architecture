package sample

import (
	"lms/internal/domain/user"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

// NewFakeUserEntity returns a new user entity with a random id, username, and email.
func NewFakeUserEntity(userType user.UserType, verified, banned bool) user.User {
	now := time.Now()
	usr := user.User{
		ID:        uuid.New(),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Type:      userType,
		Verified:  verified,
		Banned:    banned,
		CreatedAt: now,
		UpdatedAt: now,
	}
	usr.SetPassword("secret")
	return usr
}

// NewFakeUsername returns a new fake username string
func NewFakeUsername() string {
	return gofakeit.Username()
}

// NewFakeUsername returns a new fake email string
func NewFakeEmail() string {
	return gofakeit.Email()
}
