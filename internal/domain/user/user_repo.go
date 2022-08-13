package user

type Repository interface {
	GetUser(id string) (*User, error)
	CreateUser(username, email, password string, userType UserType) error
	DeleteUser(id string) error
}
