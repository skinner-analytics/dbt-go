package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func CheckAndResolveConflicts() error {
	diffCmd := exec.Command("git", "diff", "--name-only", "--diff-filter=U")
	diffOutput, err := diffCmd.Output()
	if err != nil {
		return fmt.Errorf("error checking for conflicts: %v", err)
	}

	conflictingFiles := strings.Split(strings.TrimSpace(string(diffOutput)), "\n")

	if len(conflictingFiles) > 0 && conflictingFiles[0] != "" {
		fmt.Println("The following files have conflicts:")
		for _, file := range conflictingFiles {
			fmt.Println("-", file)
		}

		for _, file := range conflictingFiles {
			prompt := promptui.Select{
				Label: fmt.Sprintf("Resolve conflict for %s", file),
				Items: []string{"Accept Incoming", "Accept Current"},
			}

			_, result, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("prompt failed: %v", err)
			}

			var resolveCmd *exec.Cmd
			if result == "Accept Incoming" {
				resolveCmd = exec.Command("git", "checkout", "--theirs", file)
			} else {
				resolveCmd = exec.Command("git", "checkout", "--ours", file)
			}

			if err := resolveCmd.Run(); err != nil {
				return fmt.Errorf("error resolving conflict for %s: %v", file, err)
			}

			addCmd := exec.Command("git", "add", file)
			if err := addCmd.Run(); err != nil {
				return fmt.Errorf("error adding resolved file %s: %v", file, err)
			}
		}
	}
	return nil
}
