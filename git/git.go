// Provides shell wrappers to work with git repository in current working directory.
package git

import (
	"sort"
	"strings"

	"github.com/gocombinator/semver"
	"github.com/gocombinator/sh"
)

// Branch returns branch name of git repository in current directory.
func Branch() string {
	return sh.Run("git", "rev-parse", "--symbolic-full-name", "--abbrev-ref", "HEAD")
}

// Status returns empty string if working tree is clean.
func Status() string {
	return sh.Run("git", "status", "--short")
}

// Clean returns true if working tree is clean, false otherwise.
func Clean() bool {
	return Status() == ""
}

// Tags returns list of tags.
func Tags() []string {
	return strings.Split(sh.Run("git", "tag"), "\n")
}

// Semvers returns list of semantic version tags.
func Semvers() []semver.Semver {
	var tags = Tags()
	var semvers = make([]semver.Semver, 0)
	for _, tag := range tags {
		var parsed = semver.Parse(tag)
		if parsed != semver.Empty {
			semvers = append(semvers, parsed)
		}
	}
	sort.Sort(sort.Reverse(semver.Slice(semvers)))
	return semvers
}

// LatestSemver returns latest semantic version tag.
func LatestSemver() semver.Semver {
	var semvers = Semvers()
	if len(semvers) == 0 {
		return semver.Empty
	}
	return semvers[0]
}
