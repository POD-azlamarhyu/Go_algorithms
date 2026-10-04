package services

import (
	"fmt"
	"go-algorithms/src"
	"time"
	"log/slog"
)

func BitSearchService(size int){
	start := time.Now()
	slog.Info("running bit search service", slog.String("time:", start.String()))
	randamData := src.RandomIntSlice(size)
	linearData := src.LinearIntSlice(size)
	target := 124
	slog.Info("list data", slog.Any("data", randamData[:size]))
	slog.Info("linear data", slog.Any("data", linearData[:size]))

	sum, subset := src.BitSearch(&linearData, target)

	if sum != nil {
		slog.Info("Found subset with",slog.String("target", fmt.Sprint(target)), slog.Any("sum", *sum), slog.Any("subset", *subset))
	} else {
		slog.Info("Not found")
	}
	endTime := time.Since(start)
	slog.Info("Bit search service completed", slog.Any("time taken", endTime))
}