/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra/blob/main/README.md

func init() {
	rootCmd.AddCommand(lsbCmd)
	lsbCmd.Flags().BoolVarP(&showAllFiles, "all", "a", false, "Show all changed files, not just .sql and .yml")
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   root.Render("dg"),
	Short: root.Render("dbt-go is a tool writing in GO to help improve the dx of dbt"),
	Long: dbt.Render(`dbt-go works to improve the DX of the dbt cli. For example:

This tool simplifies complex command-line operations by packaging 
long commands into easy-to-use shortcuts. You can run the same 
complex tasks with simple commands.`),
}
