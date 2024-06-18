package main

import (
	"fmt"
)

func task4_1() {
	count := 0
	for i := 100000; i <= 999999; i++ {
		c := i
		d := 0
		for c > 0 {
			d = d*10 + c%10
			c /= 10
		}
		if d == i {count++}
	}

	fmt.Println("Количество зеркальных билетов среди всех шестизначных номеров от 100000 до 999999:", count)
}

func task4_2() {
	var width, height int
	fmt.Print("Ширина: ")
	fmt.Scan(&width)
	fmt.Print("Высота: ")
	fmt.Scan(&height)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if (i+j)%2 == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}

func task4_3() {
	var height int
	fmt.Print("Высота: ")
	fmt.Scan(&height)

	for i := 0; i < height; i++ {
		for j := 0; j < height-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	task4_3()
}