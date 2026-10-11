package src

import (
	"fmt"
	"log/slog"
)

func BitSearch[T Data](data *[]T, target T)(*T,*[]T){
	n := len(*data)
	slog.Info("Starting bit search", slog.Any("data length", n))
	var rsum *T = nil
	var rsubset *[]T = nil
	slog.Info("Starting bit search loop", slog.Any("rsum", rsum), slog.Any("rsubset", rsubset))
	slog.Info("Before entering loop", slog.Any("bit",1 << n))
	for bit := 0; bit < (1 << n); bit++ {
		var sum T
		fmt.Println("Current bit:", bit)
		subset := make([]T, 0)
		for i := 0; i < n; i++{
			fmt.Println("Current i:", i, "Current bit mask:", 1 << i, "Current bit value:", bit & (1 << i))
			if bit & (1 << i) != 0{
				sum += (*data)[i]
				subset = append(subset,(*data)[i])
			}
		}
		fmt.Println("Current sum:", sum, "Current subset:", subset)
		if sum == target {
			rsum = &sum
			rsubset = &subset
			break
		}
	}
	slog.Info("Finished bit search loop", slog.Any("rsum", *rsum), slog.Any("rsubset", *rsubset))
	return rsum,rsubset
}