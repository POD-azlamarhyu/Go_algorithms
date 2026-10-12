package services

import (
	"go-algorithms/src"
	"log/slog"
)

func DepthFirstSearchService() {
	slog.Info("Starting Depth First Search Service", slog.String("module", "depth_first_search_service"))
	graph := src.NewGraph(4)
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 2)
	graph.AddEdge(2, 3)

	visited := make(map[int]bool)
	graph.DepthFirstSearch(0, visited)
	slog.Info("Finished Depth First Search Service", slog.String("module", "depth_first_search_service"))
}