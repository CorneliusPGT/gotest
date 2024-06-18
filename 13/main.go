package main

import (
	"fmt"
)

func mergeA(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}

func t13_1() {
	arr1 := []int{1, 3, 5, 7}
	arr2 := []int{2, 4, 6, 8, 10}

	mergedArray := mergeA(arr1, arr2)
	fmt.Println( mergedArray)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func t13_2() {
	arr := []int{5, 1, 4, 2, 8, 3}

	fmt.Println("Исходный массив:", arr)
	bubbleSort(arr)
	fmt.Println(arr)
}

func main() {
	t13_2()
}
