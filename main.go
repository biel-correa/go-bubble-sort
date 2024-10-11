package main

import (
	"fmt"
	"math/rand/v2"
)

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}

func main() {
	array := make([]int, 10)
	for i := range array {
		array[i] = rand.IntN(100)
	}
	fmt.Println("Unsorted array:", array)
	fmt.Println("Sorted array:", BubbleSort(array))
}
