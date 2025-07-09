package analyzer

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"strings"
)

type AuthorStats struct {
	Name         string
	Email        string
	LinesAdded   int
	LinesDeleted int
	Commits      int
	Files        map[string]struct{}
}

func AnalyzeRepo(path string, extFilter string) (map[string]*AuthorStats, error) {
	stats := make(map[string]*AuthorStats)

	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	cIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		return nil, err
	}

	err = cIter.ForEach(func(c *object.Commit) error {
		author := c.Author.Email
		if _, ok := stats[author]; !ok {
			stats[author] = &AuthorStats{
				Name:  c.Author.Name,
				Email: author,
				Files: make(map[string]struct{}),
			}
		}
		stats[author].Commits++

		// Parse patch to count lines
		if c.NumParents() > 0 {
			parent, _ := c.Parents().Next()
			patch, _ := parent.Patch(c)
			for _, fileStat := range patch.Stats() {
				if extFilter != "" && !strings.HasSuffix(fileStat.Name, extFilter) {
					continue 
				}

				stats[author].LinesAdded += fileStat.Addition
				stats[author].LinesDeleted += fileStat.Deletion
				stats[author].Files[fileStat.Name] = struct{}{}
			}
		}

		return nil
	})

	return stats, err
}
