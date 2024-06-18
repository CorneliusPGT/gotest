package main

import (
	"fmt"
	"math/rand"
)

func isEven(num int) bool {
	return num%2 == 0
}

func getPoints() (int, int, int, int, int, int) {
	x1 := rand.Intn(21) - 10
	y1 := rand.Intn(21) - 10
	x2 := rand.Intn(21) - 10
	y2 := rand.Intn(21) - 10
	x3 := rand.Intn(21) - 10
	y3 := rand.Intn(21) - 10
	return x1, y1, x2, y2, x3, y3
}

func transformPoints(x, y int) (int, int) {
	newX := 2*x + 10
	newY := -3*y - 5
	return newX, newY
}

func t11_1() {

	num := rand.Intn(100) + 1
	if isEven(num) {
		fmt.Printf("%d - четное число", num)
	} else {
		fmt.Printf("%d - нечетное число", num)
	}
}

func t11_2() {
	x1, y1, x2, y2, x3, y3 := getPoints()

	newX1, newY1 := transformPoints(x1, y1)
	newX2, newY2 := transformPoints(x2, y2)
	newX3, newY3 := transformPoints(x3, y3)

	fmt.Printf("Было (%d, %d), стало (%d, %d)\n", x1, y1, newX1, newY1)
	fmt.Printf("Было (%d, %d), стало (%d, %d)\n", x2, y2, newX2, newY2)
	fmt.Printf("Было (%d, %d), стало (%d, %d)\n", x3, y3, newX3, newY3)
}

func multiply(num int) (result int) {
	result = num * 2
	return
}

func addTen(num int) (result int) {
	result = num + 10
	return
}

func t11_3() {
	num := 5
	fmt.Println(num)
	num = multiply(num)
	fmt.Println(num)
	num = addTen(num)
	fmt.Println(num)
}

var global1 = 10
var global2 = 20
var global3 = 30


func add1(num int) int {
	return num + global1
}

func add2(num int) int {
	return num + global2
}

func add3(num int) int {
	return num + global3
}

func t11_4() {
	num := 5
	fmt.Printf("Исходное число: %d\n", num)
	num = add1(num)
	fmt.Printf("После добавления global1: %d\n", num)
	num = add2(num)
	fmt.Printf("После добавления global2: %d\n", num)
	num = add3(num)
	fmt.Printf("После добавления global3: %d\n", num)

}

func main() {
	t11_4()
}
