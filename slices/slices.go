package main

import "fmt"

func SliceOfSlice(s []int) {
	// slicing the slice creates a new slice
	// but does not copy the array data
	s = s[2:4]
	// modifying the sub-slice
	// changes the array of slice in main function as well
	for i := range s {
		s[i] += 10
	}
	fmt.Println("f1", s, len(s), cap(s))
}

func CopyOfSliceOfSlice(s []int) {
	s = s[2:4]
	s2 := make([]int, len(s))
	copy(s2, s)

	// or if you prefer less efficient, but more concise version:
	// s2 := append([]int{}, s[2:4]...)

	for i := range s2 {
		s2[i] += 10
	}

	fmt.Println("f1", s2, len(s2), cap(s2))
}

func GrowingSlices(l, c int) ([]int, []int) {
	// make a slice with length 3 and capacity 4
	s := make([]int, l, c)

	// initialize to 1,2,3
	s[0] = 1
	s[1] = 2
	s[2] = 3

	// capacity of the array is 4
	// adding one more number fits in the initial array
	s2 := append(s, 4)

	// modify the elements of the array
	// s and s2 still share the same array
	for i := range s2 {
		s2[i] += 10
	}

	// this append grows the array past its capacity
	// new array must be allocated for s3
	s3 := append(s2, 5)

	// modify the elements of the array to see the result
	for i := range s3 {
		s3[i] += 10
	}

	return s2, s3
}
