/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"os"
)

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
