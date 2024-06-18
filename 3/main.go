package main

import (
	"fmt"
)

func task3_1() {
	var number int

	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	for i := 0; i <= number; i++ {
		fmt.Println(i)
	}
}

func task3_2() {
	var num1 int
	var num2 int

	fmt.Print("Первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Второе число: ")
	fmt.Scan(&num2)

	for i := 0; i <= num1+num2; i++ {

		fmt.Println(i)
	}
}

func task3_3() {

	var price, discountP float64

	fmt.Print("Цена товара: ")
	fmt.Scan(&price)
	fmt.Print("Скидка: ")
	fmt.Scan(&discountP)

	if discountP > 30 {
		discountP = 30
	}

	discount := price * discountP / 100
	if discount > 2000 {
		discount = 2000
	}

	fmt.Println("Сумма скидки:", discount)

}

func task3_4() {
	var a, b, c int

	for {
		if a < 10 {
			a++
		}
		if b < 100 {
			b++
		}
		if c < 1000 {
			c++
		}
		if a == 10 && b == 100 && c == 1000 {
			break
		}
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func main() {
	task3_4()
}
