package main

import (
	"fmt"
	"math"		
)

func t7_1() {
	var x float64
	var n int

	fmt.Println("Введите значение x:")
	fmt.Scan(&x)

	fmt.Println("Введите количество знаков после запятой для точности:")
	fmt.Scan(&n)
	epsilon := 1.0 / math.Pow(10, float64(n))

	ex := 1.0
	term := 1.0

	for i := 1; math.Abs(term) >= epsilon; i++ {
		term *= x / float64(i)
		ex += term
	}

	fmt.Printf("Результат e^%.2f с точностью %d знаков после запятой: %.15f\n", x, n, ex)
}

func t7_2() {
	var deposit float64
	var year float64
	var years int

	fmt.Println("Введите сумму вклада:")
	fmt.Scan(&deposit)

	fmt.Println("Введите годовую процентную ставку:")
	fmt.Scan(&year)

	fmt.Println("Введите количество лет вклада:")
	fmt.Scan(&years)

	month := year / 100 / 12

	totalMonths := years * 12

	amount := deposit
	var bank float64

	for i := 0; i < totalMonths; i++ {
		interest := amount * month
		amount += interest
		roundedAmount := float64(int64(amount*100)) / 100
		bank += amount - roundedAmount
		amount = roundedAmount
	}

	fmt.Printf("Сумма на руки: %.2f\n", amount)
	fmt.Printf("Сумма зачисленная в пользу банка: %.15f\n", bank)
}


func main() {
	t7_2()
}