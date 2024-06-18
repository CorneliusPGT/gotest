package students

import (
	"fmt"
	"strconv"
)

type Student struct {
	Name  string
	Age   int
	Grade int
}

func NewStudent(name, ageStr, gradeStr string) (*Student, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return nil, fmt.Errorf("некорректно: %v", err)
	}

	grade, err := strconv.Atoi(gradeStr)
	if err != nil {
		return nil, fmt.Errorf("некорректно: %v", err)
	}

	return &Student{Name: name, Age: age, Grade: grade}, nil
}