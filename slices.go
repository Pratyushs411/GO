package main

import "fmt"

var arr = [5]int{10, 20, 30, 40, 50}
var slice = arr[1:4] // slice from index 1 to 3 (excludes index 4)

func PrintSlice(slice []int) {
	fmt.Println("Slice:", slice)
}
func modifySlice(slice []int) {
	slice[0] = 99
	fmt.Println("After modifying slice[0] = 99")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}

var s = make([]int, 3) // prefered as capacity can be defined
// When you append() past the capacity, Go allocates a new array, which can be costly.
// when you use " : " cap(slice) = len(arr) - startIndex
func appendSlice(s []int) {
	s = append(s, 100)
	s = append(s, 200)
	fmt.Println("After appending:", s)
}

func slicing(s []int) {
	sub := s[1:4] // slice from index 1 to 3
	fmt.Println("Sub-slice (s[1:4]):", sub)
}

func lengthOfSlice(s []int) {
	fmt.Println("Length:", len(s))
}

func capacityOfSlice(s []int) {
	fmt.Println("Capacity:", cap(s))
}

var src = []int{1, 2, 3}

func copySlice(src []int) []int {
	dest := make([]int, len(src))
	copy(dest, src)
	return dest
}

// Nil and empty slices
var nilSlice []int
var emptySlice = []int{}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)
	}
	return matrix
}
