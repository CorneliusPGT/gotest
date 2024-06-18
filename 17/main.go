package main

import (
	"fmt"
	"strings"
)

func splitN(numbers []int) ([]int, []int) {
	var evens []int
	var odds []int
	for _, number := range numbers {
		if number%2 == 0 {
			evens = append(evens, number)
		} else {
			odds = append(odds, number)
		}
	}
	return evens, odds
}

func t17_1() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens, odds := splitN(numbers)

	fmt.Println("Чётные числа:", evens)
	fmt.Println("Нечётные числа:", odds)
}

func parseTest(sentences []string, chars []rune) [][]int {
	result := make([][]int, len(sentences))

	for i, sentence := range sentences {
		words := strings.Fields(sentence)
		if len(words) == 0 {
			result[i] = make([]int, len(chars))
			for j := range chars {
				result[i][j] = -1
			}
			continue
		}
		lastWord := words[len(words)-1]
		result[i] = make([]int, len(chars))

		for j, char := range chars {
			index := strings.IndexRune(lastWord, char)
			result[i][j] = index
		}
	}
	return result
}

func t17_2() {
	sentences := []string{"Hello world", "Hello Skillbox", "Привет Мир", "Привет Skillbox"}
	chars := []rune{'H', 'E', 'L', 'П', 'М'}
	indices := parseTest(sentences, chars)
	for i, sentence := range sentences {
		fmt.Printf("Предложение: %s\n", sentence)
		for j, char := range chars {
			fmt.Printf("Символ '%c' на позиции %d в последнем слове\n", char, indices[i][j])
		}
	}
}

func main() {
	t17_2()
}
