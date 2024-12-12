package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func MergeChanges() error {
	mergeCmd := exec.Command("git", "merge", "origin/main", "--no-commit")
	mergeOutput, err := mergeCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(mergeOutput), "CONFLICT") {
			fmt.Println("Merge conflicts detected. Please resolve them manually.")
		} else {
			return fmt.Errorf("error merging changes: %v\n%s", err, mergeOutput)
		}
	}
	return nil
}
