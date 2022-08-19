package memory

import (
	"lms/internal/domain/user"
	"sync"

	"github.com/google/uuid"
)

// UserRepo fulfills the Repository interface of user entity.
type UserRepo struct {
	users map[uuid.UUID]*user.User
	mutex sync.RWMutex
}

// NewUserRepo is a factory function to generate a new repository of users
func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[uuid.UUID]*user.User),
	}
}

// Get finds a user by ID
func (ur *UserRepo) Get(id uuid.UUID) (*user.User, error) {
	ur.mutex.RLock()
	defer ur.mutex.RUnlock()
	if usr, ok := ur.users[id]; ok {
		return usr, nil
	}
	return nil, user.ErrUserNotFound
}

// GetByUsername finds a user by username
func (ur *UserRepo) GetByUsername(username string) (*user.User, error) {
	ur.mutex.RLock()
	defer ur.mutex.RUnlock()
	for _, usr := range ur.users {
		if usr.Username == username {
			return usr, nil
		}
	}
	return nil, user.ErrUserNotFound
}

// Create will add a new user to the repository
func (ur *UserRepo) Create(username, email, password string, userType user.UserType) (*user.User, error) {
	//TODO: Make sure user isn't already in repository
	ur.mutex.Lock()
	defer ur.mutex.Unlock()

	newUser := &user.User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	err := newUser.SetPassword(password)
	if err != nil {
		return nil, err
	}
	// Prepare the user for create (Fill other properties...)
	newUser.PrepareForCreate()
	ur.users[newUser.ID] = newUser
	return newUser, nil
}

// Delete removes a user from the memory repository
func (ur *UserRepo) Delete(id uuid.UUID) error {
	usr, err := ur.Get(id)
	if err != nil {
		return err
	}
	delete(ur.users, usr.ID)
	return nil
}
