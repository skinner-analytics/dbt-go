package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetChangesCount() (string, error) {
	compareCmd := exec.Command("git", "rev-list", "HEAD...origin/main", "--count")
	compareOutput, err := compareCmd.Output()
	if err != nil {
		return "", fmt.Errorf("error comparing branches: %v", err)
	}
	return strings.TrimSpace(string(compareOutput)), nil
}
