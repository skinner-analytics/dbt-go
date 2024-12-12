package git

import (
	"fmt"
	"os/exec"
)

func CommitChanges() error {
	commitCmd := exec.Command("git", "commit", "-m", "Resolve merge conflicts and merge changes from main")
	commitOutput, err := commitCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error committing changes: %v\n%s", err, commitOutput)
	}
	fmt.Println("Successfully resolved conflicts and committed changes.")
	return nil
}
