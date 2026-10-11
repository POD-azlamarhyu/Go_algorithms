package main

import (
	"go-algorithms/src/services"
	"log/slog"
	"os"
)

const commandListLength int = 2
const bitDataSize int = 61
const listDataSize int = 1_000_000_000

const (
	bitSearchCommand string = "bit"
	linearSearchCommand string = "linear"
	permutationCommand string = "perm"
)

func main(){
	slog.Info("アルゴリズムのサンプルを実行する CLI", slog.String("module", "main"))
	args := os.Args

	if len(args) < commandListLength {
		slog.Warn("Usage: algorithms <command>")
		slog.Warn("help - Show this help messages")
		return
	}


	switch args[1] {
		case bitSearchCommand:
			services.BitSearchService(bitDataSize)
		case linearSearchCommand:
			services.LinearSearchService(listDataSize)
		case permutationCommand:
			services.PermutationSearchService(listDataSize)
		default:
			slog.Warn("Unknown command:",slog.String("inputed command:", args[1]))
			return
	}
	slog.Info("Go cli finished.")
}