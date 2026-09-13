package src

import (
	"log/slog"
)

func BitSearch[T Data](data *[]T, target T)(*T,*[]T){
	n := len(*data)
	slog.Info("Starting bit search", slog.Any("data length", n))
	var rsum *T = nil
	var rsubset *[]T = nil

	for bit := 0; bit < (1 << n); bit++ {
		var sum T
		subset := make([]T, 0)
		for i := 0; i < n; i++{
			if (bit >> i) & 1 == 1 {
				sum += (*data)[i]
				subset = append(subset,(*data)[i])
			}
		}
		if sum == target {
			rsum = &sum
			rsubset = &subset
			break
		}
	}

	return rsum,rsubset
}