package main

import (
	"fmt"
	"math"
)

func task2_1() {
	var x, y int

	fmt.Println("Определение координатной плоскости точки.")
	fmt.Print("Введите X: ")
	fmt.Scan(&x)
	fmt.Print("Введите Y: ")
	fmt.Scan(&y)

	if x > 0 && y > 0 {
		fmt.Println("Точка находится в первой четверти.")
	} else if x < 0 && y > 0 {
		fmt.Println("Точка находится во второй четверти.")
	} else if x < 0 && y < 0 {
		fmt.Println("Точка находится в третьей четверти.")
	} else if x > 0 && y < 0 {
		fmt.Println("Точка находится в четвертой четверти.")
	} else if x == 0 && y == 0 {
		fmt.Println("Точка находится в начале координат.")
	} else if x == 0 {
		fmt.Println("Точка находится на оси Y.")
	} else {
		fmt.Println("Точка находится на оси X.")
	}
}

func task2_2() {
	var num1, num2, num3 int

	fmt.Println("Проверка, что одно из чисел — положительное.")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&num3)

	if num1 > 0 || num2 > 0 || num3 > 0 {
		fmt.Println("Хотя бы одно из чисел положительное.")
	} else {
		fmt.Println("Ни одно из чисел не является положительным.")
	}
}

func task2_3() {
	var num1, num2, num3 int

	fmt.Println("Проверка на совпадение чисел.")
	fmt.Print("Введите первое число: ")
	fmt.Scan(&num1)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&num2)
	fmt.Print("Введите третье число: ")
	fmt.Scan(&num3)

	if num1 == num2 || num1 == num3 || num2 == num3 {
		fmt.Println("Есть совпадающие числа.")
	} else {
		fmt.Println("Совпадающих чисел нет.")
	}
}

func task2_4() {
	var cost, coin1, coin2, coin3 int

	fmt.Println("Сумма без сдачи.")
	fmt.Print("Введите стоимость товара: ")
	fmt.Scan(&cost)
	fmt.Print("Введите номинал первой монеты: ")
	fmt.Scan(&coin1)
	fmt.Print("Введите номинал второй монеты: ")
	fmt.Scan(&coin2)
	fmt.Print("Введите номинал третьей монеты: ")
	fmt.Scan(&coin3)

	if coin1 == cost || coin2 == cost || coin3 == cost ||
		coin1+coin2 == cost || coin1+coin3 == cost || coin2+coin3 == cost ||
		coin1+coin2+coin3 == cost {
		fmt.Println("Можно заплатить без сдачи.")
	} else {
		fmt.Println("Нельзя заплатить без сдачи.")
	}
}

func task2_5() {
	var a, b, c float64

	fmt.Println("Решение квадратного уравнения.")
	fmt.Print("Введите коэффициент a: ")
	fmt.Scan(&a)
	fmt.Print("Введите коэффициент b: ")
	fmt.Scan(&b)
	fmt.Print("Введите коэффициент c: ")
	fmt.Scan(&c)

	discriminant := b*b - 4*a*c

	if discriminant > 0 {
		root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Println("Уравнение имеет два корня:", root1, "и", root2)
	} else if discriminant == 0 {
		root := -b / (2 * a)
		fmt.Println("Уравнение имеет один корень:", root)
	} else {
		fmt.Println("Уравнение не имеет действительных корней.")
	}
}

func task2_6() {
	var ticket int

	fmt.Println("Проверка билета.")
	fmt.Print("Введите четырёхзначный номер билета: ")
	fmt.Scan(&ticket)

	d1 := ticket / 1000
	d2 := (ticket / 100) % 10
	d3 := (ticket / 10) % 10
	d4 := ticket % 10

	if d1 == d4 && d2 == d3 {
		fmt.Println(ticket, "-> зеркальный билет")
	} else if d1+d2 == d3+d4 {
		fmt.Println(ticket, "-> счастливый билет")
	} else {
		fmt.Println(ticket, "-> обычный билет")
	}
}

func task2_7() {
	var guess, response int

	fmt.Println("Игра 'Угадай число'.")
	fmt.Println("Загадайте число от 1 до 10.")

	low := 1
	high := 10
	attempts := 0
	maxAttempts := 4

	for attempts < maxAttempts {
		guess = (low + high) / 2
		fmt.Printf("Ваше число ", guess, " ? (1 - угадали, 2 - больше, 3 - меньше): ")
		fmt.Scan(&response)

		if response == 1 {
			fmt.Println("Компьютер угадал число!")
			return
		} else if response == 2 {
			low = guess + 1
		} else if response == 3 {
			high = guess - 1
		}
		attempts++
	}

	fmt.Println("Компьютер не смог угадать число.")
}

func main() {
	task2_7()
	fmt.Println()

}
