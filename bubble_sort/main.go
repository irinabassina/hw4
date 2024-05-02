package main

import "fmt"

func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{42, 7, 13, 15, 1, 60}

	fmt.Println("Исходный массив:", arr)
	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
