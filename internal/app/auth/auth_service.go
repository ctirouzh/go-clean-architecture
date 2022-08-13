package auth

import (
	"lms/internal/domain/user"
)

type AuthService struct {
	userRepo user.UserRepository
}

func NewAuthService(userRepo user.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (service *AuthService) SignUp(username, email, password string) error {
	return service.userRepo.CreateUser(username, email, password)
}
