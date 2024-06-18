package main

import (
	"fmt"
		
)

func t5_1() {
	var month string
	fmt.Print("Введите месяц: ")
	fmt.Scan(&month)

	switch month {
	case "декабрь", "январь", "февраль":
		fmt.Println("Зима.")
	case "март", "апрель", "май":
		fmt.Println("Весна.")
	case "июнь", "июль", "август":
		fmt.Println("Лето.")
	case "сентябрь", "октябрь", "ноябрь":
		fmt.Println("Осень.")
	default:
		fmt.Println("Неверный месяц")
	}
}

func t5_2() {
	var day string

	fmt.Print("День (пн, вт, ср, чт, пт): ")
	fmt.Scan(&day)

	switch day {
	case "пн":
		fmt.Println("понедельник")
		fallthrough
	case "вт":
		fmt.Println("вторник")
		fallthrough
	case "ср":
		fmt.Println("среда")
		fallthrough
	case "чт":
		fmt.Println("четверг")
		fallthrough
	case "пт":
		fmt.Println("пятница")
	default:
		if day != "пн" && day != "вт" && day != "ср" && day != "чт" && day != "пт" {
			fmt.Println("Неверный день недели.")
		}
	}
}


func main() {
	t5_2()
}