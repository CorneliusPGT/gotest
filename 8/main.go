package main

import (
	"fmt"
	"strings"
	"unicode"
	"strconv"
)

func t8_1() {
	text := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	fmt.Println("Определение количества слов, начинающихся с большой буквы в строке:")
	fmt.Println(text)
	words := strings.Fields(text)
	count := 0
	for _, word := range words {
		if len(word) > 0 && unicode.IsUpper(int32(word[0])) {
			count++
		}
	}
	fmt.Printf("%d слов с большой буквы", count)
}

func t8_2() {
	text := "a10 10 20b 20 30c30 30 dd"
	fmt.Println("Исходная строка:")
	fmt.Println(text)
	var numbers []string
	words := strings.Fields(text)
	for _, word := range words {
		if num, err := strconv.Atoi(word); err == nil {
			numbers = append(numbers, strconv.Itoa(num))
		}
	}
	fmt.Println("В строке содержатся числа в десятичном формате:")
	fmt.Println(strings.Join(numbers, " "))

}

func main() {
	t8_2()
}