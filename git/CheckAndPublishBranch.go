package git

import (
	"dg/style"
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func CheckAndPublishBranch(currentBranch string) error {
	branchCheckCmd := exec.Command("git", "rev-parse", "--symbolic-full-name", "--abbrev-ref", "@{u}")
	_, err := branchCheckCmd.Output()
	if err != nil {
		prompt := promptui.Prompt{
			Label:     "Branch is not published. Would you like to publish it to the remote?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf(style.Red.Render("Prompt failed: %v"), err)
		}

		if strings.ToLower(result) == "y" {
			pushCmd := exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
			pushOutput, err := pushCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error pushing changes: %v\n%s", err, pushOutput)
			}
			fmt.Println("Successfully pushed the changes to the remote.")
		}
	} else {
		pushCmd := exec.Command("git", "push")
		pushOutput, err := pushCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error pushing changes: %v\n%s", err, pushOutput)
		}
		fmt.Println("Successfully pushed the changes to the remote.")
	}
	return nil
}
