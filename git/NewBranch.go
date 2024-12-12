package git

import (
	"dg/style"
	"fmt"
	"os/exec"
	"strings"
)

func NewBranch(branchName string) (string, error) {
	output, err := exec.Command("git", "switch", "-c", branchName).Output()
	if err != nil {
		return "", fmt.Errorf(style.Red.Render("Error creating a new branch: %v"), err)
	}
	return strings.TrimSpace(string(output)), nil
}
