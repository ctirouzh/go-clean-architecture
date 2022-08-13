package user

import "github.com/tahadostifam/go-clean-architecture/app/entity"

type UserService struct {
	userRepo entity.UserRepository
}

func NewUserService(userRepo entity.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) GetUser(id string) (*entity.User, error) {
	return service.userRepo.GetUser(id)
}
func (service *UserService) CreateUser(username, email, password string) error {
	return service.userRepo.CreateUser(username, email, password)
}
func (service *UserService) DeleteUser(id string) error {
	return service.userRepo.DeleteUser(id)
}
