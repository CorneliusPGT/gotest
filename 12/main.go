package main

import (
	"fmt"
)

func t12_1() {
	var nums [10]int
	evenCount := 0
	oddCount := 0
	fmt.Println("Введите 10 целых чисел:")
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
		if nums[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Printf("Чётных %d, нечётных %d\n", evenCount, oddCount)
}

func reverseArray(arr [10]int) [10]int {
	reversed := [10]int{}
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		reversed[i] = arr[j]
	}
	return reversed
}

func t12_2() {
	var nums [10]int
	fmt.Println("Введите 10 целых чисел:")
	for i := 0; i < len(nums); i++ {
		fmt.Scan(&nums[i])
	}
	reversed := reverseArray(nums)
	fmt.Println(reversed)
}

func main() {
	t12_2()
}
