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
	if err := repo.db.Where("id=?", id.String()).Find(&usr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user.ErrUserNotFound
		}
		return nil, err
	}
	return usr, nil
}

// GetByUsername finds a user by username
func (repo *UserRepo) GetByUsername(username string) (*user.User, error) {
	return nil, errors.New("GetByUsername is not implemented")
}

// Create will add a new user to the postgres repository
func (repo *UserRepo) Create(username, email, password string, userType user.UserType) (*user.User, error) {
	return nil, errors.New("Create is not unimplemented")
}

// Delete removes a user from the postgres repository
func (repo *UserRepo) Delete(id uuid.UUID) error {
	return errors.New("Delete is not unimplemented")
}
