/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

////////////////////////
//////// styles ////////

// https://github.com/charmbracelet/lipgloss/blob/master/README.md

var dbt = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF694A")).
	PaddingTop(1).
	PaddingLeft(2).
	Width(120)

//////// styles ////////
////////////////////////

func init() {
	rootCmd.AddCommand(lsbCmd)
	lsbCmd.Flags().BoolVarP(&showAllFiles, "all", "a", false, "Show all changed files, not just .sql and .yml")
}

///////////////////////////
//////// commands ////////

var rootCmd = &cobra.Command{
	Use:   "dg",
	Short: "dbt-go is a tool writting in GO to help improve the dx of dbt",
	Long: `This works to improve the dx of the dbt cli. For example:

This tool simplifies complex command-line operations by packaging 
long commands into easy-to-use shortcuts. You can run the same 
complex tasks with simple commands.`,
}

var lsbCmd = &cobra.Command{
	Use:   "lsb",
	Short: "List changed files on the current branch",
	Long:  `By default, lists only changed .sql and .yml files. Use --all to show all changed files.`,
	RunE:  runLsb,
}

var showAllFiles bool

//////// commands ////////
///////////////////////////

///////////////////////////
//////// functions ////////

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func runLsb(cmd *cobra.Command, args []string) error {
	output, err := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main").Output()
	if err != nil {
		return fmt.Errorf("Error executing git diff: %v", err)
	}

	files := strings.Split(string(output), "\n")
	sqlOrYamlFound := false
	filesFound := false

	for _, file := range files {
		if file == "" {
			continue
		}

		isSqlOrYaml := strings.HasSuffix(file, ".sql") || strings.HasSuffix(file, ".yml")

		if showAllFiles || isSqlOrYaml {
			if isSqlOrYaml {
				cmd.Println(dbt.Render(file))
				sqlOrYamlFound = true
			} else if showAllFiles {
				cmd.Println(dbt.Render(file)) // Ensure consistent formatting
			}
			filesFound = true
		}
	}

	if !filesFound {
		cmd.Println(dbt.Render("No files were changed."))
	} else if !showAllFiles && !sqlOrYamlFound {
		cmd.Println(dbt.Render("No .sql or .yml files were changed."))
	}

	return nil
}

//////// functions ////////
///////////////////////////
