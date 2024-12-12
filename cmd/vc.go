package cmd

import (
	"dg/style"
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
		fmt.Printf("\n---\nYou chose %s!\n", style.Orange.Render(m.choice))
	}
	return nil
}
