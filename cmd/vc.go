package cmd

import (
	"dg/git"
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
	p := tea.NewProgram(vcModel{})

	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	if m, ok := m.(vcModel); ok && m.choice != "" {
		switch m.choice {
		case "Create A New Branch":
			output, err := git.NewBranch(m.branchName)
			if err != nil {
				return fmt.Errorf("error creating new branch: %v", err)
			}
			fmt.Println(output)
		case "Select An Existing Branch":
			// Handle selecting an existing branch
		case "Exit":
			// Handle exit
		}
	}
	return nil
}
