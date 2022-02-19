package main

import "fmt"

const size1 = 4
const size2 = 5
const sizeResult = size1 + size2

func mergeSortedArrays(a [size1]int, b [size2]int) [sizeResult]int {
	var result [sizeResult]int
	j, k := 0, 0
	for i := 0; i < sizeResult; i++ {
		if j == size1 {
			result[i] = b[k]
			k++
			continue
		} else if k == size2 {
			result[i] = a[j]
			continue
		}
		if a[j] < b[k] {
			result[i] = a[j]
			j++
		} else {
			result[i] = b[k]
			k++
		}
	}

	return result
}

func main() {
	firstArray := [size1]int{1, 3, 4, 9}
	secondArray := [size2]int{2, 5, 6, 7, 8}
	fmt.Println(mergeSortedArrays(firstArray, secondArray))
}
