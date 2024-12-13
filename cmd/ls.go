package cmd

import (
	"dg/git"
	"dg/style"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(lsCmd)
	lsCmd.Flags().BoolVarP(&showAllFiles, "all", "a", false, "Show all changed files, not just .sql and .yml")
}

var lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List changed files on the current branch",
	Long:  `By default, lists only changed .sql and .yml files. Use --all to show all changed files.`,
	RunE:  runls,
}

var showAllFiles bool

func runls(cmd *cobra.Command, args []string) error {
	files, err := git.GitDiff()
	if err != nil {
		return fmt.Errorf("error getting changed files: %v", err)
	}

	sqlOrYamlFound := false
	filesFound := false

	for _, file := range files {
		if file == "" {
			continue
		}

		isSqlOrYaml := strings.HasSuffix(file, ".sql") || strings.HasSuffix(file, ".yml")

		if showAllFiles || isSqlOrYaml {
			if isSqlOrYaml {
				cmd.Println(style.DgIndent.Render(file))
				sqlOrYamlFound = true
			} else if showAllFiles {
				cmd.Println(style.DgIndent.Render(file))
			}
			filesFound = true
		}
	}

	if !filesFound {
		cmd.Println(style.DgIndent.Render("No files were changed."))
	} else if !showAllFiles && !sqlOrYamlFound {
		cmd.Println(style.DgIndent.Render("No .sql or .yml files were changed."))
	}

	return nil
}
