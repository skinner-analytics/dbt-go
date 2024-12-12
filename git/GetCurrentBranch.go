package git

import (
	"dg/style"
	"fmt"
	"os/exec"
	"strings"
)

func GetCurrentBranch() (string, error) {
	output, err := exec.Command("git", "rev-parse", "--symbolic-full-name", "--abbrev-ref", "@{u}").Output()
	if err != nil {
		return "", fmt.Errorf(style.Red.Render("Error getting current branch: %v"), err)
	}
	return strings.TrimSpace(string(output)), nil
}
