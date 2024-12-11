/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsbCmd)
	lsbCmd.Flags().BoolVarP(&showAllFiles, "all", "a", false, "Show all changed files, not just .sql and .yml")
}

var lsbCmd = &cobra.Command{
	Use:   orange.Render("lsb"),
	Short: "List changed files on the current branch",
	Long:  `By default, lists only changed .sql and .yml files. Use --all to show all changed files.`,
	RunE:  runLsb,
}

var showAllFiles bool

func runLsb(cmd *cobra.Command, args []string) error {
	output, err := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main").Output()
	if err != nil {
		return fmt.Errorf("error executing git diff: %v", err)
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
				cmd.Println(orange_indent.Render(file))
				sqlOrYamlFound = true
			} else if showAllFiles {
				cmd.Println(orange_indent.Render(file))
			}
			filesFound = true
		}
	}

	if !filesFound {
		cmd.Println(orange_indent.Render("No files were changed."))
	} else if !showAllFiles && !sqlOrYamlFound {
		cmd.Println(orange_indent.Render("No .sql or .yml files were changed."))
	}

	return nil
}
