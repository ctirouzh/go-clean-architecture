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

type MockRepository struct {
	users map[uuid.UUID]*User
}

func NewMockRepository(users map[uuid.UUID]*User) *MockRepository {
	return &MockRepository{users: users}
}

func (repo *MockRepository) Get(id uuid.UUID) (*User, error) {
	usr, exists := repo.users[id]
	if !exists {
		return nil, ErrUserNotFound
	}
	return usr, nil
}

func (repo *MockRepository) Create(username, email, password string, userType UserType) (*User, error) {

	for _, user := range repo.users {
		if user.Username == username {
			return nil, ErrUsernameAlreadyTaken
		}
		if user.Email == email {
			return nil, ErrEmailAlreadyTaken
		}
	}
	usr := &User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	if err := usr.SetPassword(password); err != nil {
		return nil, err
	}
	usr.PrepareForCreate()
	repo.users[usr.ID] = usr
	return usr, nil
}

func (repo *MockRepository) Delete(id uuid.UUID) error {
	usr, err := repo.Get(id)
	if err != nil {
		return err
	}
	delete(repo.users, usr.ID)
	return nil
}
