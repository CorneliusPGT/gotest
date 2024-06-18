package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Student struct {
	name  string
	age   int
	grade int
}

func main() {
	students := make(map[string]*Student)
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студентов:")

	for scanner.Scan() {
		line := scanner.Text()

		if line == "EOF" {
			break
		}

		fields := strings.Fields(line)
		if len(fields) < 3 {
			fmt.Println("Некорректный формат строки")
			continue
		}

		name := fields[0]
		age := fields[1]
		grade := fields[2]

		student := newStudent(name, age, grade)
		students[name] = student
	}

	fmt.Println("Студенты из хранилища:")
	for _, student := range students {
		fmt.Printf("%s %d %d\n", student.name, student.age, student.grade)
	}
}

func newStudent(name, ageStr, gradeStr string) *Student {
	age := convertToInt(ageStr)
	grade := convertToInt(gradeStr)
	return &Student{name: name, age: age, grade: grade}
}

func convertToInt(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Printf("Ошибка при конвертации строки '%s' в число: %v\n", str, err)
		return 0
	}
	return result
}