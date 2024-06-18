package main

import (
	"fmt"
)

func calculate(x int16, y uint8, z float32) float32 {
	S := 2*float32(x) + float32(y)*float32(y) - 3/z
	return S
}

func t15_1() {
	x := int16(5)
	y := uint8(3)
	z := float32(2.0)

	result := calculate(x, y, z)
	fmt.Println(result)
}

func execute(f func(int, int) int, a int, b int) {
	defer func() {
		result := f(a, b)
		fmt.Println(result)
	}()
}

func t15_2() {
	execute(func(a int, b int) int {
		return a + b
	}, 3, 5)
	execute(func(a int, b int) int {
		return a * b
	}, 3, 5)
	execute(func(a int, b int) int {
		return a - b
	}, 10, 5)
}

func main() {
	t15_2()
}