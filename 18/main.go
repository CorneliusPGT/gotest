package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func t18_1() {
	arr := []int{29, 10, 14, 37, 14, 3, 10, 25, 42, 1}
	fmt.Println(arr)
	insertionSort(arr)
	fmt.Println(arr)
}

func t18_2() {
	bubbleSortReverse := func(arr []int) {
		n := len(arr)
		for i := 0; i < n-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if arr[j] < arr[j+1] {
					arr[j], arr[j+1] = arr[j+1], arr[j]
				}
			}
		}
	}
	arr := []int{29, 10, 14, 37, 14, 3, 10, 25, 42, 1}
	fmt.Println(arr)
	bubbleSortReverse(arr)
	fmt.Println(arr)
}

func main() {
	t18_2()
}
