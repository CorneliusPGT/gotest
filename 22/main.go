package main

import (
	"bufio"
	"fmt"
	"os"
	"student_app/students"
)

func main() {
	storage := students.NewStorage()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введите данные студентов:")

	for scanner.Scan() {
		line := scanner.Text()

		if line == "EOF" {
			break
		}

		student, err := students.ParseLine(line)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		storage.AddStudent(student)
	}

	fmt.Println("Студенты из хранилища:")
	for _, student := range storage.GetAllStudents() {
		fmt.Printf("%s %d %d\n", student.Name, student.Age, student.Grade)
	}
}
