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
