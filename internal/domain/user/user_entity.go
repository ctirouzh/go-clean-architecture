package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyUserPassword = errors.New("password cannot be empty")
	ErrUserNotVerified   = errors.New("user not verified")
	ErrBannedUser        = errors.New("user is banned")
)

type User struct {
	ID           uuid.UUID
	Username     string
	Email        string
	Type         UserType
	passwordHash string
	Verified     bool
	Banned       bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user User) IsAdmin() bool {
	return user.Type == USER_TYPE_ADMIN
}

func (user User) IsTeacher() bool {
	return user.Type == USER_TYPE_TEACHER
}

func (user User) IsStudent() bool {
	return user.Type == USER_TYPE_STUDENT
}

func (user User) IsVerified() bool {
	return user.Verified
}

func (user User) IsBanned() bool {
	return user.Banned
}

func (user *User) SetPassword(password string) error {
	if len(password) == 0 {
		return ErrEmptyUserPassword
	}
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.passwordHash = string(b)
	return nil
}

func (user User) IsPasswordVerified(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.passwordHash), []byte(password)) == nil
}

func (user *User) PrepareForCreate() {
	user.ID = uuid.New()
	user.Verified = false
	user.Banned = false
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
}
