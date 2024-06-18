package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Необходимо указать хотя бы один файл")
		return
	}
	firstFileContent, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println("Ошибка при чтении файла", args[0], err)
		return
	}
	if len(args) == 1 {
		fmt.Println(string(firstFileContent))
	} else {
		secondFileContent, err := os.ReadFile(args[1])
		if err != nil {
			fmt.Println("Ошибка при чтении файла", args[1], err)
			return
		}
		combinedContent := strings.Join([]string{string(firstFileContent), string(secondFileContent)}, "\n")
		fmt.Println(combinedContent)
	}
}
