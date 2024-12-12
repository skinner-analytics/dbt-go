package git

import (
	"os/exec"
	"strings"
)

func GitDiff() ([]string, error) {
	output, err := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main").Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(string(output), "\n")
	return files, nil
}
