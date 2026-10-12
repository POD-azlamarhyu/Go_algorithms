package services
import (
	"fmt"
	"go-algorithms/src"
	"log/slog"
	"time"
)
func BinarySearchService(size int){
	start := time.Now()
	slog.Info("running binary search service", slog.String("time:", start.String()))
	linearData := src.LinearIntSlice(size)
	target := 56
	slog.Info("linear data", slog.Any("data", linearData[:size]))

	result := src.BinarySearch(&linearData, target)

	if result != nil {
		slog.Info("Found target", slog.String("target", fmt.Sprint(target)), slog.Any("result", *result))
	} else {
		slog.Info("Not found")
	}
	endTime := time.Since(start)
	slog.Info("Binary search service completed", slog.Any("time taken", endTime))
}