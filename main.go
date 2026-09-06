package main

import (
	"fmt"
	"go-algorithms/src"
	"time"
)

func main(){
	fmt.Println("Searching for 3 in a random slice of 100000000 integers")
	randomDataSize := 1000000000
	start := time.Now()

	fmt.Println("Generating random data...")
	data := src.RandomIntSlice(randomDataSize)
	fmt.Println(data[:100])
	result, idx := src.LinearSearch(&data, 3)
	if result != nil && idx != nil {
		fmt.Println("Found:", *result, "at index:", *idx)
	} else {
		fmt.Println("Not found")
	}
	endTime := time.Since(start)
	fmt.Println("Time taken:", endTime)
}