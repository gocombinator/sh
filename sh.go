package sh

import (
	"os/exec"
	"strings"
)

// MaybeRun executes command returning stdout on success.
func MaybeRun(name string, args ...string) (string, error) {
	if outBytes, err := exec.Command(name, args...).Output(); err != nil {
		return "", err
	} else {
		return strings.Trim(string(outBytes), "\n\r "), nil
	}
}

// Run executes command returning whitespace trimmed stdout.
// Returns empty string on error.
func Run(name string, args ...string) string {
	if out, err := MaybeRun(name, args...); err != nil {
		return ""
	} else {
		return out
	}
}
