package auth

import (
	"github.com/tahadostifam/go-clean-architecture/app/entity"
)

type AuthService struct {
	userRepo entity.UserRepository
}

func NewAuthService(userRepo entity.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (service *AuthService) SignUp(username, email, password string) error {
	return service.userRepo.CreateUser(username, email, password)
}
