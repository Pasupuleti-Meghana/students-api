package storage
import "Pasupuleti-Meghana/students-api/internal/types"

type Storage interface {
	CreateStudent(name string, email string , age string) (int64, error)
	GetStudentById(id int64) (types.Student, error)
}