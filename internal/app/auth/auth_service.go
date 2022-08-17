package auth

import (
	"errors"
	"lms/internal/domain/user"
)

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserNotVerified = errors.New("user not verified")
	ErrBannedUser      = errors.New("user is banned")
)

// Service is the auth usecase service struct
type Service struct {
	userRepo   user.Repository
	jwtManager *JwtManager
}

// NewService returns a pointer to auth service
func NewService(userRepo user.Repository, jwtManager *JwtManager) *Service {
	return &Service{userRepo: userRepo, jwtManager: jwtManager}
}

// SignUp creates a new student type user in user repository
func (service *Service) SignUp(username, email, password string) (*user.User, error) {
	// business rule: only students can sign up...
	return service.userRepo.Create(username, email, password, user.USER_TYPE_STUDENT)
}

// SignIn tries to sign in user with given username and password,
// and returns an access token on success.
func (service *Service) SignIn(username, password string) (string, error) {
	usr, findErr := service.userRepo.GetByUsername(username)
	if findErr != nil {
		return "", findErr
	}
	if !usr.IsVerified() {
		return "", ErrUserNotVerified
	}
	if usr.IsBanned() {
		return "", ErrBannedUser
	}
	if !usr.IsPasswordVerified(password) {
		return "", ErrInvalidPassword
	}
	token, genErr := service.jwtManager.Generate(usr)
	if genErr != nil {
		return "", genErr
	}
	return token, nil
}
