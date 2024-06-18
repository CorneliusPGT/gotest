package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func t9_1() {
	file, err := os.OpenFile("messages.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()
	fmt.Println("Введите строки для записи в файл. Для завершения введите 'exit'.")

	for {
		var text string
		fmt.Print("> ")
		fmt.Scanln(&text)

		if text == "exit" {
			break
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		line := fmt.Sprintf(timestamp, text)

		_, err := file.WriteString(line)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}
	}

	fmt.Println("Программа завершена.")
}

func t9_2() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Не удалось получить информацию о файле:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Println("Файл пуст.")
		return
	}

	buffer := make([]byte, fileInfo.Size())

	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	content := string(buffer)
	fmt.Println("Содержимое файла:")
	fmt.Println(content)
}

func t9_3() {
	file, err := os.OpenFile("readonly.txt", os.O_CREATE|os.O_RDONLY, 0444)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	file.Close()

	file, err = os.OpenFile("readonly.txt", os.O_WRONLY, 0444)
	if err != nil {
		fmt.Println("Ошибка при открытии файла для записи:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Попытка записи.")
	if err != nil {
		fmt.Println("Запись в файл не удалась:", err)
	} else {
		fmt.Println("Запись в файл прошла успешно.")
	}
}

func t9_4() {
	filename := "messages_ioutil.txt"


	fmt.Println("Введите строки для записи в файл. Для завершения введите 'exit'.")

	for {
		var text string
		fmt.Print("> ")
		fmt.Scanln(&text)

		if text == "exit" {
			break
		}

		timestamp := time.Now().Format("2006-01-02 15:04:05")
		line := fmt.Sprintf(timestamp, text)
		content, err := ioutil.ReadFile(filename)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Ошибка при чтении файла:", err)
			return
		}

		newContent := append(content, line...)

		err = ioutil.WriteFile(filename, newContent, 0644)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}

	}

	fmt.Println("Программа завершена.")
}

func main() {
	t9_4()
}
