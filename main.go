package main

import (
	"fmt"
	"go-algorithms/src/services"
)

func main(){
	fmt.Println("Searching for 3 in a random slice integers")
	randomDataSize := 1000000000
	bitDataSize := 61
	services.LinearSearchService(randomDataSize)
	services.BitSearchService(bitDataSize)
}