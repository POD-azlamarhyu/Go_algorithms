package services

import (
	"go-algorithms/src"
	"time"
	"log/slog"
)

func PermutationSearchService(size int){
	start := time.Now()
	slog.Info("running linear search service",slog.String("time:", start.String()))
	randamData := src.MultiDimenstionSlice(size)
	slog.Info("list data", slog.Any("data", randamData))

	target := 100

	dist, route := src.PermutationSearch(*randamData,target)

	if dist != nil && route != nil {
		slog.Info("Found:", slog.Any("Distance",*dist), slog.Any("short route", *route))
	} else {
		slog.Info("Not found")
	}
	endTime := time.Since(start)
	slog.Info("Linear search service completed", slog.Any("time taken", endTime))
	
}