package githubapi

import (
	"fmt"
	"strings"
)

type Commits struct {
	CommitList []Commit
}

// Return a string of comma-separated commit hashes from the available array of commits
func (c *Commits) GetCommitsAsList() string {
	var s []string
	for _, v := range c.CommitList {
		s = append(s, v.Sha)
	}
	return fmt.Sprint(strings.Join(s, " "))
}