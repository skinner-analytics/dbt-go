/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dg",
	Short: "dbt-go is a tool writting in GO to help improve the dx of the dbt ADLC",
	Long: `This works to improve the dbt ADLC. For example:

This tool simplifies complex command-line operations by packaging 
long commands into easy-to-use shortcuts. You can run the same 
complex tasks with just a simple command.`,
}

var lsbCmd = &cobra.Command{
	Use:   "lsb",
	Short: "List changed files on the current branch",
	Long:  `By default, lists only changed .sql and .yml files. Use --all to show all changed files.`,
	RunE:  runLsb,
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run dbt commands",
	Long:  `Run dbt commands.`,
	RunE:  runRun,
}

var showAllFiles bool

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(lsbCmd)
	lsbCmd.Flags().BoolVarP(&showAllFiles, "all", "a", false, "Show all changed files, not just .sql and .yml")
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
			cmd.Println(file)
			filesFound = true
			if isSqlOrYaml {
				sqlOrYamlFound = true
			}
		}
	}

	if !filesFound {
		cmd.Println("No files were changed.")
	} else if !showAllFiles && !sqlOrYamlFound {
		cmd.Println("No .sql or .yml files were changed.")
	}

	return nil
}

func runRun(cmd *cobra.Command, args []string) error {
	fmt.Println("Running dbt commands")
	return nil
}
