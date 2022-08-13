package user

import (
	"time"

	"github.com/google/uuid"
)

type UserType int

const (
	USER_TYPE_ADMIN UserType = iota + 1
	USER_TYPE_TEACHER
	USER_TYPE_STUDENT
)

type UserRepository interface {
	GetUser(id string) (*User, error)
	CreateUser(username, email, password string) error
	DeleteUser(id string) error
}

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	RoleType     UserType  `json:"role_type"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (user User) IsAdmin() bool {
	return user.RoleType == USER_TYPE_ADMIN
}

func (user User) IsTeacher() bool {
	return user.RoleType == USER_TYPE_TEACHER
}

func (user User) IsStudent() bool {
	return user.RoleType == USER_TYPE_STUDENT
}

func (user *User) PrepareForCreate() {
	user.ID = uuid.New().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) PrepareForUpdate() {
	user.UpdatedAt = time.Now()
}
