package main

import (
	"fmt"
)

func task1_1() {
	var first int
	var second int
	var third int
	pass := 275
	fmt.Println("Введите результат первого экзамена: ")
	fmt.Scan(&first)
	fmt.Println("Введите результат второго экзамена: ")
	fmt.Scan(&second)
	fmt.Println("Введите результат третьего экзамена: ")
	fmt.Scan(&third)
	total := first + second + third
	fmt.Println("Количество набранных баллов: ", total)
	if total < pass {
		fmt.Print("Вы не поступили")
	} else {
		fmt.Print("Вы поступили")
	}
}

func task1_2() {
	var first int
	var second int
	var third int
	fmt.Println("Введите первое число: ")
	fmt.Scan(&first)
	fmt.Println("Введите второе число: ")
	fmt.Scan(&second)
	fmt.Println("Введите третье число: ")
	fmt.Scan(&third)
	if first > 5 || second > 5 || third > 5 {
		fmt.Println("Есть число больше 5")
	} else {
		fmt.Println("Нет числа больше 5")
	}
}

func task1_3() {
	var amount int
	max := 100000

	fmt.Println("Банкомат.")
	fmt.Print("Введите сумму снятия со счёта: ")
	fmt.Scan(&amount)

	if amount%100 != 0 {
		fmt.Println("Сумма должна быть кратна 100.")
	} else if amount > max {
		fmt.Println("Максимальная сумма снятия", max, "рублей")
	} else {
		fmt.Println("Операция успешно выполнена.")
		fmt.Println("Вы сняли", amount, "рублей")
	}

}

func task1_4() {
	var num1, num2, num3 int
	count := 0

	fmt.Println("Три числа.")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&num3)

	if num1 >= 5 {
		count++
	}
	if num2 >= 5 {
		count++
	}
	if num3 >= 5 {
		count++
	}

	fmt.Println("Среди введённых чисел ", count, " больше или равны 5")
}

func task1_5() {
	var day, guests int
	var billAmount, discount, serviceCharge, totalAmount float64

	fmt.Print("Введите день недели (1-Понедельник, ..., 7-Воскресенье): ")
	fmt.Scan(&day)
	fmt.Print("Введите число гостей: ")
	fmt.Scan(&guests)
	fmt.Print("Введите сумму чека: ")
	fmt.Scan(&billAmount)

	totalAmount = billAmount

	if day == 1 {
		discount = 0.10 * billAmount
		totalAmount -= discount
		fmt.Println("Скидка по понедельникам:", discount)
	}

	if day == 5 && billAmount > 10000 {
		extraDiscount := 0.05 * billAmount
		discount += extraDiscount
		totalAmount -= extraDiscount
		fmt.Println("Скидка по пятницам:", extraDiscount)
	}

	if guests > 5 {
		serviceCharge = 0.10 * billAmount
		totalAmount += serviceCharge
		fmt.Println("Надбавка на обслуживание:", serviceCharge)
	}

	fmt.Println("Сумма к оплате:", totalAmount)
}

func task1_6() {
	var numStudents, numGroups, studentNumber int

	fmt.Print("Введите общее количество студентов: ")
	fmt.Scan(&numStudents)
	fmt.Print("Введите количество групп: ")
	fmt.Scan(&numGroups)
	fmt.Print("Введите порядковый номер студента: ")
	fmt.Scan(&studentNumber)

	groupNumber := (studentNumber-1)%numGroups + 1

	fmt.Println("Студент с номером", studentNumber, "будет в группе", groupNumber)
}

func main() {
	task1_6()
	fmt.Println()

}
