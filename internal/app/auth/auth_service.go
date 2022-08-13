package auth

import (
	"lms/internal/domain/user"
)

// Service is the auth usecase service struct
type Service struct {
	userRepo user.Repository
}

// NewService returns a pointer to auth service
func NewService(userRepo user.Repository) *Service {
	return &Service{userRepo: userRepo}
}

// SignUp creates a new student type user in user repository
func (service *Service) SignUp(username, email, password string) error {
	// business rule: only students can sign up...
	return service.userRepo.Create(username, email, password, user.USER_TYPE_STUDENT)
}
