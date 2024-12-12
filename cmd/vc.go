package cmd

import (
	"dg/git"
	"dg/style"
	"fmt"

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
	currentBranch, err := git.GetCurrentBranch()
	if err != nil {
		return err
	}

	fmt.Printf("%s %s\n", style.Orange.Render("Current branch:"), currentBranch)

	if err := git.FetchChanges(); err != nil {
		return err
	}

	changesCount, err := git.GetChangesCount()
	if err != nil {
		return err
	}

	if changesCount == "0" {
		fmt.Println("No changes to pull from main.")
		return nil
	}

	if err := git.MergeChanges(); err != nil {
		return err
	}

	if err := git.CheckAndResolveConflicts(); err != nil {
		return err
	}

	if err := git.CommitChanges(); err != nil {
		return err
	}

	if err := git.CheckAndPublishBranch(); err != nil {
		return err
	}

	return nil
}
