package main 

import "fmt"

func t10_1(a int, b int){
	if a > b {
		a, b = b, a
	}

	sum := 0
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			sum += i
		}
	}

	fmt.Println("Сумма чётных чисел в диапазоне:", sum)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func t10_2(x int, y int){

	fmt.Println("До:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)

	swap(&x, &y)

	fmt.Println("После:")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
}

func main() {
t10_1(1, 10)
t10_2(3, 7)
}