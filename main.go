package main

import (
	"flag"
	"log"

	"github.com/gorillassi/gitAnalys/analyzer"
	"github.com/gorillassi/gitAnalys/output"
)

func main() {
	ext := flag.String("ext", "", "Filter files by extension (e.g. .go)")
	sortBy := flag.String("sort", "lines", "Sort by: lines | commits | name")
	flag.Parse()

	repoPath := "."
	if flag.NArg() > 0 {
		repoPath = flag.Arg(0)
	}

	stats, err := analyzer.AnalyzeRepo(repoPath, *ext)
	if err != nil {
		log.Fatal(err)
	}

	output.PrintStats(stats, *sortBy)
}
