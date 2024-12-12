package cmd

import (
	"dg/style"
	"os"

	"github.com/spf13/cobra"
)

// https://github.com/spf13/cobra/blob/main/README.md

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   style.Orange.Render("dg"),
	Short: style.Orange.Render("\ndg is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
	Long:  style.Orange.Render("dg is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
}
