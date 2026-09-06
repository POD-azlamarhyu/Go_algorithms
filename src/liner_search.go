package src

import (
	"fmt"
)

type Data interface{
	~int | ~float64 | ~string
}

func LinearSearch[T Data](data *[]T, target T) (*T, *T) {
	d := *data
	var r *T = nil
	var idx *T = nil
	fmt.Println("Searching for", target, "in data of size", len(d))
	for i := 0; i < len(d); i++ {
		if d[i] == target {
			r = &d[i]
			break
		}
	}
	return r, idx
}