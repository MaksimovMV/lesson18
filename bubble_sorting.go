package main

import "fmt"

const size = 6

func bubbleSort(a [size]int) [size]int {
	for i := 0; i < size-1; i++ {
		changes := false
		for j := 0; j < size-1-i; j++ {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
				changes = true
			}
		}
		if !changes {
			break
		}
	}
	return a
}

func main() {
	array := [size]int{1, 5, 8, 3, 2, 15}
	fmt.Println(bubbleSort(array))
}
