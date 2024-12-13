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
	Use:   style.Dg.Render("dg"),
	Short: style.Dg.Render("\ndg is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
	Long:  style.Dg.Render("dg is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
}
