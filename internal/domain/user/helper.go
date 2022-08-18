package user

import (
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/google/uuid"
)

// NewUser returns a random user, with an ID, a fake username, and email. The returned
// user is not verified (verified: false) and permitted (banned: false). Its password and
// timestamps are set to "secret", and time.Now() respectively.
func NewUser(userType UserType) User {
	//TODO: check if userType is valid
	now := time.Now()
	usr := User{
		ID:        uuid.New(),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Type:      userType,
		CreatedAt: now,
		UpdatedAt: now,
	}
	usr.SetPassword("secret")
	return usr
}

// NewVerifiedUser returns a random verified user, with an ID, a fake username, and email.
// The returned user is permitted (banned: false). Its password and timestamps are set to
// "secret", and time.Now() respectively.
func NewVerifiedUser(userType UserType) User {
	//TODO: check if userType is valid
	now := time.Now()
	usr := User{
		ID:        uuid.New(),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Type:      userType,
		Verified:  true,
		Banned:    false,
		CreatedAt: now,
		UpdatedAt: now,
	}
	usr.SetPassword("secret")
	return usr
}

// NewBannedUser returns a random banned user, with an ID, a fake username, and email.
// The returned user is verified (verified: true). Its password and timestamps are set to
// "secret", and time.Now() respectively.
func NewBannedUser(userType UserType) User {
	//TODO: check if userType is valid
	now := time.Now()
	usr := User{
		ID:        uuid.New(),
		Username:  gofakeit.Username(),
		Email:     gofakeit.Email(),
		Type:      userType,
		Verified:  true,
		Banned:    true,
		CreatedAt: now,
		UpdatedAt: now,
	}
	usr.SetPassword("secret")
	return usr
}
