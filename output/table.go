package output

import (
	"fmt"
	"github.com/gorillassi/gitAnalys/analyzer"
	"sort"
	"strings"
)

func PrintStats(stats map[string]*analyzer.AuthorStats) {
	headers := []string{"Author", "Commits", "Lines Added", "Lines Deleted", "Files Changed"}

	rows := [][]string{}
	for _, s := range stats {
		rows = append(rows, []string{
			s.Name,
			fmt.Sprintf("%d", s.Commits),
			fmt.Sprintf("%d", s.LinesAdded),
			fmt.Sprintf("%d", s.LinesDeleted),
			fmt.Sprintf("%d", len(s.Files)),
		})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][2] > rows[j][2]
	})

	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, col := range row {
			if len(col) > widths[i] {
				widths[i] = len(col)
			}
		}
	}

	printRow := func(row []string) {
		for i, col := range row {
			fmt.Printf("%-*s  ", widths[i], col)
		}
		fmt.Println()
	}

	printRow(headers)
	fmt.Println(strings.Repeat("-", sum(widths)+len(widths)*2))

	for _, row := range rows {
		printRow(row)
	}
}

func sum(xs []int) int {
	s := 0
	for _, x := range xs {
		s += x
	}
	return s
}
