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