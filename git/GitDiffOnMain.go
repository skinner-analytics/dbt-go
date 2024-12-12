package git

import (
	"os/exec"
	"strings"
)

func GitDiffOnMain() ([]string, error) {
	output, err := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main").Output()
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(output), "\n")

	return result, nil
}
