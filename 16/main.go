package main

import (
	"fmt"
	"math/rand"
)

func t16_1() {

	array := make([]int, 20)
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}

	fmt.Println("Массив:", array)
	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)
	count := 0
	found := false
	for i := 0; i < len(array); i++ {
		if array[i] == num {
			count = len(array) - i - 1
			found = true
			break
		}
	}

	if found {
		fmt.Printf("Чисел после %d: %d", num, count)
	} else {
		fmt.Println("Число не найдено в массиве.")
	}
}

func t16_2() {
	array := []int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array)

	var num int
	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	index := -1
	for i := 0; i < len(array); i++ {
		if array[i] == num {
			index = i
			break
		}
	}

	if index != -1 {
		fmt.Println("Первое вхождение по индексу", index)
	} else {
		fmt.Println("Число не найдено")
	}
}

func main() {
	t16_2()
}
