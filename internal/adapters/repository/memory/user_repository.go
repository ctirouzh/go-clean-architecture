package memory

import (
	"lms/internal/domain/user"
	"sync"

	"github.com/google/uuid"
)

// UserRepo fulfills the Repository interface of user entity.
type UserRepo struct {
	users map[uuid.UUID]user.User
	*sync.Mutex
}

// NewUserRepo is a factory function to generate a new repository of users
func NewUserRepo() *UserRepo {
	return &UserRepo{
		users: make(map[uuid.UUID]user.User),
	}
}

// Get finds a user by ID
func (ur UserRepo) Get(id uuid.UUID) (user.User, error) {
	if usr, ok := ur.users[id]; ok {
		return usr, nil
	}
	return user.User{}, user.ErrUserNotFound
}

// Create will add a new user to the repository
func (ur UserRepo) Create(username, email, password string, userType user.UserType) error {
	// Make sure user isn't already in repository
	for _, usr := range ur.users {
		if usr.Username == username {
			return user.ErrUsernameAlreadyTaken
		}
		if usr.Email == email {
			return user.ErrEmailAlreadyTaken
		}
	}
	ur.Lock()
	defer ur.Unlock()
	user := user.User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	// Prepare the user for create (fills other properties...)
	user.PrepareForCreate()
	ur.users[user.ID] = user
	return nil
}

func (ur UserRepo) Delete(id uuid.UUID) error {
	usr, err := ur.Get(id)
	if err != nil {
		return err
	}
	delete(ur.users, usr.ID)
	return nil
}
