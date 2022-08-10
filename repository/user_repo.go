package repository

import (
	"database/sql"
	"errors"

	"github.com/tahadostifam/go-clean-architecture/app/entity"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) GetUser(id string) (*entity.User, error) {

	var user entity.User
	query := "SELECT * FROM users WHERE id = ?"
	if err := ur.db.QueryRow(query, id).Scan(&user); err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, errors.New("record not found")
		default:
			return nil, err
		}
	}

	return &user, nil //
}

func (ur *UserRepository) CreateUser(username, email, password string) error {
	var user entity.User
	// TODO: Check if user with given username or email exists
	q := "INSERT INTO users(id,username,email,role_type,password_hash,created_at,updated_at) VALUES(?, ?, ?, ?, ?, ?, ?)"
	stmt, prepareErr := ur.db.Prepare(q)
	if prepareErr != nil {
		return prepareErr
	}
	defer stmt.Close()

	if _, execErr := stmt.Exec(
		user.ID,
		user.Username,
		user.Email,
		user.RoleType,
		user.PasswordHash,
		user.CreatedAt,
		user.UpdatedAt,
	); execErr != nil {
		return execErr
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id string) error {

	user, getErr := ur.GetUser(id)
	if getErr != nil {
		return getErr
	}
	if user == nil {
		// TODO: handle this type of errors by an specific user repository error type
		return errors.New("user not found")
	}
	// TODO: prepare a query to delete the user
	return nil
}
