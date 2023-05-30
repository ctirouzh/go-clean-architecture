package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyUserPassword = errors.New("password cannot be empty")
)

type User struct {
	ID           uuid.UUID `gorm:"primaryKey"`
	Username     string    `gorm:"uniqueIndex;not null;size:32;default:null"`
	Email        string    `gorm:"uniqueIndex;not null;size:42;default:null"`
	Type         UserType  `gorm:"index;size:1;default:3"`
	PasswordHash string    `gorm:"not null;size:64;default:null"`
	Verified     bool      `gorm:"default:false"`
	Banned       bool      `gorm:"default:false"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (user *User) IsAdmin() bool {
	return user.Type == USER_TYPE_ADMIN
}

func (user *User) IsTeacher() bool {
	return user.Type == USER_TYPE_TEACHER
}

func (user *User) IsStudent() bool {
	return user.Type == USER_TYPE_STUDENT
}

func (user *User) IsVerified() bool {
	return user.Verified
}

func (user *User) IsBanned() bool {
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
	user.PasswordHash = string(b)
	return nil
}

func (user *User) IsPasswordVerified(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil
}

func (user *User) PrepareForCreate() {
	user.ID = uuid.New()
	user.Verified = false
	user.Banned = false
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now
}
