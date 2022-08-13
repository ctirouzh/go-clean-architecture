package user

import "lms/internal/domain/user"

type UserService struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) GetUser(id string) (*user.User, error) {
	return service.userRepo.GetUser(id)
}
func (service *UserService) CreateUser(username, email, password string) error {
	return service.userRepo.CreateUser(username, email, password)
}
func (service *UserService) DeleteUser(id string) error {
	return service.userRepo.DeleteUser(id)
}
