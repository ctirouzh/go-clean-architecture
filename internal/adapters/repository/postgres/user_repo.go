package postgres

import (
	"errors"
	"lms/internal/domain/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UserRepo fulfills the Repository interface of user entity.
type UserRepo struct {
	db *gorm.DB
}

// NewUserRepo is a factory function to generate a new postgres repository of users
func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

// Get finds a user by ID
func (repo *UserRepo) Get(id uuid.UUID) (*user.User, error) {
	usr := &user.User{}
	if err := repo.db.Where("id=?", id.String()).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return usr, nil
}

// GetByUsername finds a user by username
func (repo *UserRepo) GetByUsername(username string) (*user.User, error) {
	usr := &user.User{}
	if err := repo.db.Where("username=?", username).First(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return usr, nil
}

// GetByUsernameOrEmail finds a user by username, or email
func (repo *UserRepo) GetByUsernameOrEmail(username, email string) (*user.User, error) {
	usr := &user.User{}
	if err := repo.db.Where("username=? OR email=?", username).First(&usr).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return usr, nil
}

// Create will add a new user to the postgres repository
func (repo *UserRepo) Create(username, email, password string, userType user.UserType) (*user.User, error) {
	if _, findErr := repo.GetByUsernameOrEmail(username, email); findErr == nil {
		return nil, user.ErrUsernameOrEmailAlreadyTaken
	}
	usr := &user.User{
		Username: username,
		Email:    email,
		Type:     userType,
	}
	err := usr.SetPassword(password)
	if err != nil {
		return nil, err
	}
	// Prepare the user for create (Fill other properties...)
	usr.PrepareForCreate()
	if createErr := repo.db.Create(&usr).Error; createErr != nil {
		return nil, createErr
	}
	return usr, nil
}

// Delete removes a user from the postgres repository
func (repo *UserRepo) Delete(id uuid.UUID) error {
	return errors.New("Delete is not unimplemented")
}
