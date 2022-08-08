package repository

import (
	"database/sql"

	"github.com/tahadostifam/go-clean-architecture/model"
)

type StudentRepository struct {
	dbClient *sql.DB
}

func NewStudentRepository(c *sql.DB) *StudentRepository {
	return &StudentRepository{
		dbClient: c,
	}
}

func (s *StudentRepository) CreateStudent(student model.Student) error {
	stmt, prepareErr := s.dbClient.Prepare("insert into student(name, last_name, age, gender) values(?, ?, ?, ?)")
	if prepareErr != nil {
		return prepareErr
	}
	defer stmt.Close()

	stmt.Exec(
		student.Name,
		student.LastName,
		student.Age,
		student.Gender,
	)

	return nil
}

func (s *StudentRepository) GetStudentByName() {}
