package git

import (
	"dg/style"
	"fmt"
	"os/exec"
	"strings"
)

func GetCurrentBranch() (string, error) {
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return "", fmt.Errorf(style.Red.Render("Error getting current branch: %v"), err)
	}
	return strings.TrimSpace(string(branchOutput)), nil
}
