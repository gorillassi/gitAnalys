package main

import (
	"log"
	"os"
	"github.com/gorillassi/gitAnalys/analyzer"
	"github.com/gorillassi/gitAnalys/output"
)

func main() {
	repoPath := "."
	if len(os.Args) > 1 {
		repoPath = os.Args[1]
	}

	stats, err := analyzer.AnalyzeRepo(repoPath)
	if err != nil {
		log.Fatal(err)
	}

	output.PrintStats(stats)
}
