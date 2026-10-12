package main

import (
	"go-algorithms/src/services"
	"log/slog"
	"os"
)

const commandListLength int = 2
const bitDataSize int = 61
const listDataSize int = 100_000_000

const (
	bitSearchCommand string = "bit"
	linearSearchCommand string = "linear"
	permutationCommand string = "perm"
	binarySearchCommand string = "binary"
	depthFirstSearchCommand string = "dfs"
)

func printUsage() {
	slog.Warn("Usage: algorithms <command>")
	slog.Warn("help - Show this help messages")
	slog.Warn("bit - Perform bit search")
	slog.Warn("linear - Perform linear search")
	slog.Warn("perm - Perform permutation search")
	slog.Warn("binary - Perform binary search")
	slog.Warn("dfs - Perform depth first search")
}

func main(){
	slog.Info("アルゴリズムのサンプルを実行する CLI", slog.String("module", "main"))
	args := os.Args

	if len(args) < commandListLength {
		printUsage()
		return
	}


	switch args[1] {
		case bitSearchCommand:
			services.BitSearchService(bitDataSize)
		case linearSearchCommand:
			services.LinearSearchService(listDataSize)
		case permutationCommand:
			services.PermutationSearchService(listDataSize)
		case binarySearchCommand:
			services.BinarySearchService(listDataSize)
		case depthFirstSearchCommand:
			services.DepthFirstSearchService()
		default:
			slog.Warn("Unknown command:",slog.String("inputed command:", args[1]))
			return
	}
	slog.Info("Go cli finished.")
	return
}