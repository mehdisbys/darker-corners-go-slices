package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSliceOfSlice(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		input    []int
		expected []int
	}{
		{
			name:     "Slicing a slice",
			comment:  "A slice of slice modifies the original array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 13, 14, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SliceOfSlice(tt.input)

			if !reflect.DeepEqual(tt.expected, tt.input) {
				t.Errorf("Expected %v but got %v\n", tt.expected, tt.input)
			}
		})
	}
}

func TestCopySliceOfSlice(t *testing.T) {
	tests := []struct {
		name     string
		comment  string
		input    []int
		expected []int
	}{
		{
			name:     "Slicing a copy of a slice",
			comment:  "A copy of slice does not modifies the original array",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CopyOfSliceOfSlice(tt.input)

			if !reflect.DeepEqual(tt.expected, tt.input) {
				t.Errorf("Expected %v but got %v\n", tt.expected, tt.input)
			}
		})
	}
}

func TestGrowingSlice(t *testing.T) {

	expectedS2 := []int{11, 12, 13, 14}
	expectedS3 := []int{21, 22, 23, 24, 15}

	s2, s3 := GrowingSlices(3, 4)

	if !reflect.DeepEqual(expectedS2, s2) {
		t.Errorf("Expected S2 to be %v but got %v\n", expectedS2, s2)
	}

	if !reflect.DeepEqual(expectedS3, s3) {
		t.Errorf("Expected S3 to be %v but got %v\n", expectedS3, s3)
	}
}

func TestNilSlice(t *testing.T) {
	var s []int // this is a nil slice
	s2 := []int{}

	if !reflect.DeepEqual(fmt.Sprintf("%p", s), "0x0") {
		t.Errorf("Expected S to be nil but got %v\n", fmt.Sprintf("%p", s))
	}

	if reflect.DeepEqual(fmt.Sprintf("%p", s), fmt.Sprintf("%p", s2)) { // checking inequality
		t.Errorf("Expected S and S2 to be different but they are equal %v\n", fmt.Sprintf("%p", s2))
	}

}
