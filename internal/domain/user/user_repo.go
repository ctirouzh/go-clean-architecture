package user

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrUserNotFound                = errors.New("user not found")
	ErrUsernameAlreadyTaken        = errors.New("username already taken")
	ErrEmailAlreadyTaken           = errors.New("email already taken")
	ErrUsernameOrEmailAlreadyTaken = errors.New("usename or email already taken")
)

type Repository interface {
	Get(id uuid.UUID) (*User, error)
	GetByUsername(username string) (*User, error)
	Create(username, email, password string, userType UserType) (*User, error)
	Delete(id uuid.UUID) error
}

type MockRepository struct {
	data  map[uuid.UUID]*User
	mutex sync.RWMutex
}

// NewUserRepo is a factory function to generate a new mock repository of users
func NewMockRepository() *MockRepository {
	return &MockRepository{data: make(map[uuid.UUID]*User)}
}

// AddUsers directly adds a slice of users to the mock repository.
// It don't use Create method, soi t can be used for test purposes which
// needs different mutable users.
func (repo *MockRepository) AddUsers(users []*User) {
	repo.mutex.Lock()
	defer repo.mutex.Unlock()
	for _, usr := range users {
		repo.data[usr.ID] = usr
	}
}

// Get finds a user by ID
func (repo *MockRepository) Get(id uuid.UUID) (*User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	if usr, exists := repo.data[id]; exists {
		return usr, nil
	}
	return nil, ErrUserNotFound
}

// GetByUsername retrieves user by its username, and return a user not found error on failure
func (repo *MockRepository) GetByUsername(username string) (*User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	for _, usr := range repo.data {
		if usr.Username == username {
			return usr, nil
		}
	}
	return nil, ErrUserNotFound
}

// GetByEmail retrieves user by its email, and return a user not found error on failure
func (repo *MockRepository) GetByEmail(email string) (*User, error) {
	repo.mutex.RLock()
	defer repo.mutex.RUnlock()

	for _, usr := range repo.data {
		if usr.Email == email {
			return usr, nil
		}
	}
	return nil, ErrUserNotFound
}

// Create will add a new user to the mock repository
func (repo *MockRepository) Create(username, email, password string, userType UserType) (*User, error) {

	usr := &User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	if err := usr.SetPassword(password); err != nil {
		return nil, err
	}

	if _, err := repo.GetByUsername(username); err == nil {
		return nil, ErrUsernameAlreadyTaken
	}
	if _, err := repo.GetByEmail(email); err == nil {
		return nil, ErrEmailAlreadyTaken
	}

	repo.mutex.Lock()
	defer repo.mutex.Unlock()

	usr.PrepareForCreate()
	repo.data[usr.ID] = usr
	return usr, nil
}

// Delete removes a user by id, and return a user not found error if user not exists
func (repo *MockRepository) Delete(id uuid.UUID) error {
	usr, err := repo.Get(id)
	if err != nil {
		return err
	}
	delete(repo.data, usr.ID)
	return nil
}
