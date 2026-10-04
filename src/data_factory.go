package src

import (
	"fmt"
	"math/rand/v2"
)

func RandomIntSlice(size int) []int {
	intRange := 10000
	fmt.Println("Generating random int slice of size:", size)
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = int(rand.IntN(intRange))
	}
	return slice
}

func LinearIntSlice(size int) []int {
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = i
	}
	return slice
}

func MultiDimenstionSlice(size int) *[][]int{
	slice := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}
	return &slice
}