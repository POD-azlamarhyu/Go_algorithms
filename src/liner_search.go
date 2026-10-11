package src

import (
	"log/slog"
)

type Data interface{
	~int | ~float64 | ~string
}

func LinearSearch[T Data](data *[]T, target T) (*T, *T) {
	d := *data
	var r *T = nil
	var idx *T = nil
	slog.Info("Starting linear search", slog.Any("target", target), slog.Any("data length", len(d)))
	for i := 0; i < len(d); i++ {
		if d[i] == target {
			r = &d[i]
			break
		}
	}
	return r, idx
}