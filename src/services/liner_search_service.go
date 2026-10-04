package services

import (
	"go-algorithms/src"
	"time"
	"log/slog"
)

func LinearSearchService(size int){
	start := time.Now()
	slog.Info("running linear search service",slog.String("time:", start.String()))
	randamData := src.RandomIntSlice(size)
	linearData := src.LinearIntSlice(size)
	slog.Info("list data", slog.Any("data", randamData[:size/1000]))
	slog.Info("linear data", slog.Any("data", linearData[:size/1000]))
	target := 121
	result, idx := src.LinearSearch(&randamData, target)

	if result != nil && idx != nil {
		slog.Info("Found:", slog.Any("result",result), slog.Any("at index", *idx))
	} else {
		slog.Info("Not found")
	}
	endTime := time.Since(start)
	slog.Info("Linear search service completed", slog.Any("time taken", endTime))
	
}