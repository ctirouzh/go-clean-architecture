package auth

import (
	"lms/internal/domain/user"
)

type AuthService struct {
	userRepo user.Repository
}

func NewAuthService(userRepo user.Repository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (service *AuthService) SignUp(username, email, password string) error {
	// This service only creates a new student user
	return service.userRepo.Create(username, email, password, user.USER_TYPE_STUDENT)
}
