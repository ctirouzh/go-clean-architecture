package entity

type Student struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
}

type StudentRepository interface {
	GetStudentByName(name string) (*Student, error)
	GetStudentByLastName(lastName string) (*Student, error)
	CreateStudent(student Student) error
}
