package sqlite

import (
	"Pasupuleti-Meghana/students-api/config"
	"Pasupuleti-Meghana/students-api/internal/types"
	"database/sql"
	"fmt"
	"log/slog"
)

type Sqlite struct {
	DB *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	//connect to the database 
	db, err := sql.Open("sqlite3", cfg.StoragePath)

	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err 
	}

	return &Sqlite{
		DB: db,
	}, nil

}

func (s *Sqlite) CreateStudent(name string, email string , age string) (int64, error) {
	stmt, err := s.DB.Prepare("INSERT INTO students(name,email,age) VALUES(?,?,?)")
	if err != nil {
		slog.Info("first err:", err.Error())
		return 0,nil 
	}

	defer stmt.Close()

	result, err := stmt.Exec(name,email,age)
	if err !=nil {
		slog.Info("second err:", err.Error())
		return 0, nil
	}

	lastId, err := result.LastInsertId()
	slog.Info("lastId:", lastId)
	if err !=nil {
		slog.Info("third err:", err.Error())
		return 0, nil 
	}

	return lastId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.DB.Prepare("SELECT id,name,email,age FROM students WHERE id = ? LIMIT 1")
	if err != nil {
		slog.Error("Getting error in preparing stmt", err)
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.ID,&student.Name,&student.Email,&student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			return types.Student{}, fmt.Errorf("No student found with the id %s", fmt.Sprint(id))
		}
		return types.Student{}, fmt.Errorf("Query Error %w",err)
	}

	return student,nil
}

