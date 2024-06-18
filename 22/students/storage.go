package students

import (
	"fmt"
	"strings"
)


type Storage struct {
	students map[string]*Student
}

func NewStorage() *Storage {
	return &Storage{students: make(map[string]*Student)}
}

func (s *Storage) AddStudent(student *Student) {
	s.students[student.Name] = student
}

func (s *Storage) GetAllStudents() []*Student {
	students := make([]*Student, 0, len(s.students))
	for _, student := range s.students {
		students = append(students, student)
	}
	return students
}

func ParseLine(line string) (*Student, error) {
	fields := strings.Fields(line)
	if len(fields) < 3 {
		return nil, fmt.Errorf("некорректный формат строки, ожидаются: имя возраст оценка")
	}

	return NewStudent(fields[0], fields[1], fields[2])
}