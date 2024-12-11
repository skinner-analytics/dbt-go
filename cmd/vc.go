/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(vcCmd)
}

var vcCmd = &cobra.Command{
	Use:   "vc",
	Short: "Interact with version control",
	Long:  `Simple interface to interact with version control.`,
	RunE:  runVc,
}

func runVc(cmd *cobra.Command, args []string) error {
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		return fmt.Errorf("error getting current branch: %v", err)
	}
	currentBranch := strings.TrimSpace(string(branchOutput))

	fmt.Printf("%s %s\n", orange.Render("Current branch:"), currentBranch)

	if currentBranch != "main" {
		prompt := promptui.Prompt{
			Label:     "Pull changes from main into this branch?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %v", err)
		}

		if strings.ToLower(result) == "y" {
			mergeBaseCmd := exec.Command("git", "merge-base", currentBranch, "main")
			mergeBaseOutput, err := mergeBaseCmd.Output()
			if err != nil {
				return fmt.Errorf("error checking merge-base: %v", err)
			}
			mergeBase := strings.TrimSpace(string(mergeBaseOutput))

			diffCmd := exec.Command("git", "diff", mergeBase, "main", "--name-only")
			diffOutput, err := diffCmd.Output()
			if err != nil {
				return fmt.Errorf("error checking for conflicts: %v", err)
			}

			conflictingFiles := strings.Split(strings.TrimSpace(string(diffOutput)), "\n")

			if len(conflictingFiles) > 0 && conflictingFiles[0] != "" {
				fmt.Println("Warning: The following files may have conflicts:")
				for _, file := range conflictingFiles {
					fmt.Println("-", file)
				}
				confirmPrompt := promptui.Prompt{
					Label:     "Proceed with pulling changes from main?",
					IsConfirm: true,
				}
				confirmResult, err := confirmPrompt.Run()
				if err != nil {
					return fmt.Errorf("confirmation prompt failed: %v", err)
				}
				if strings.ToLower(confirmResult) != "y" {
					fmt.Println("Pull operation cancelled.")
					return nil
				}
			}

			pullCmd := exec.Command("git", "pull", "origin", "main")
			pullOutput, err := pullCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error pulling changes: %v\n%s", err, pullOutput)
			}
			fmt.Println("Successfully pulled changes from main.")
		}
	}

	return nil
}
