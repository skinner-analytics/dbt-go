package cmd

import (
	"dg/style"
	"fmt"
	"os"
	"strings"

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
	Short: style.Orange.Render("dbt-go is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
	Long: func() string {
		asciiArt := style.GetASCIIArt()
		description := style.Orange.Render("dbt-go is a cli written in GO to help improve the DX of Analytics Engineers using dbt")
		copyright := "Copyright © 2024 Matthew Skinner"
		contact := "matthew@skinnerdev.com"

		width := 80

		centeredDescription := style.CenterText(description, width)
		centeredCopyright := style.CenterText(copyright, width)
		centeredContact := style.CenterText(contact, width)

		lines := strings.Split(asciiArt, "\n")
		centeredLines := make([]string, len(lines))
		for i, line := range lines {
			centeredLines[i] = style.CenterText(line, width)
		}

		centeredAsciiArt := strings.Join(centeredLines, "\n")

		return fmt.Sprintf(`
%s

%s

%s
%s
`, centeredDescription, centeredAsciiArt, centeredCopyright, centeredContact)
	}(),
}
