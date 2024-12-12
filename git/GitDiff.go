package git

import (
	"os/exec"
	"strings"
)

func GitDiff() ([]string, error) {
	uncommittedOutput, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return nil, err
	}

	result := strings.Split(string(uncommittedOutput), "\n")

	return result, nil
}
