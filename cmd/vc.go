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

	// Fetch the latest changes from the remote repository
	fetchCmd := exec.Command("git", "fetch")
	if err := fetchCmd.Run(); err != nil {
		return fmt.Errorf("error fetching changes: %v", err)
	}

	// Check if there are any changes on the remote main branch
	compareCmd := exec.Command("git", "rev-list", "HEAD...origin/main", "--count")
	compareOutput, err := compareCmd.Output()
	if err != nil {
		return fmt.Errorf("error comparing branches: %v", err)
	}
	changesCount := strings.TrimSpace(string(compareOutput))

	if changesCount == "0" {
		fmt.Println("No changes to pull from main.")
		return nil
	}

	// Perform the merge without committing
	mergeCmd := exec.Command("git", "merge", "origin/main", "--no-commit")
	mergeOutput, err := mergeCmd.CombinedOutput()
	if err != nil {
		if strings.Contains(string(mergeOutput), "CONFLICT") {
			fmt.Println("Merge conflicts detected. Please resolve them manually.")
		} else {
			return fmt.Errorf("error merging changes: %v\n%s", err, mergeOutput)
		}
	}

	// Check for conflicts
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

		commitCmd := exec.Command("git", "commit", "-m", "Resolve merge conflicts and merge changes from main")
		commitOutput, err := commitCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error committing changes: %v\n%s", err, commitOutput)
		}
		fmt.Println("Successfully resolved conflicts and committed changes.")
	} else {
		// Commit the changes if there are no conflicts
		commitCmd := exec.Command("git", "commit", "-m", "Merge changes from main")
		commitOutput, err := commitCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error committing changes: %v\n%s", err, commitOutput)
		}
		fmt.Println("Successfully pulled and committed changes from main.")
	}

	// Check if the branch is already published
	branchCheckCmd := exec.Command("git", "rev-parse", "--symbolic-full-name", "--abbrev-ref", "@{u}")
	_, err = branchCheckCmd.Output()
	if err != nil {
		// Branch is not published, ask the user if they want to publish it
		prompt := promptui.Prompt{
			Label:     "Branch is not published. Would you like to publish it to the remote?",
			IsConfirm: true,
		}

		result, err := prompt.Run()
		if err != nil {
			return fmt.Errorf("prompt failed: %v", err)
		}

		if strings.ToLower(result) == "y" {
			pushCmd := exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
			pushOutput, err := pushCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error pushing branch: %v\n%s", err, pushOutput)
			}
			fmt.Println("Successfully published the branch to the remote.")
		} else {
			fmt.Println("Branch not published.")
		}
	} else {
		// Branch is already published, push the changes
		pushCmd := exec.Command("git", "push")
		pushOutput, err := pushCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error pushing changes: %v\n%s", err, pushOutput)
		}
		fmt.Println("Successfully pushed the changes to the remote.")
	}

	return nil
}
