/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import "github.com/spf13/cobra"

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
