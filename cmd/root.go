/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
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
	Use:   root.Render("dg"),
	Short: root.Render("dbt-go is a cli written in GO to help improve the DX of Analytics Engineers using dbt"),
	Long: func() string {
		asciiArt := getASCIIArt()
		description := root.Render("dbt-go is a cli written in GO to help improve the DX of Analytics Engineers using dbt")
		copyright := "Copyright © 2024 Matthew Skinner"
		contact := "matthew@skinnerdev.com"

		width := 80

		centeredDescription := centerText(description, width)
		centeredCopyright := centerText(copyright, width)
		centeredContact := centerText(contact, width)

		lines := strings.Split(asciiArt, "\n")
		centeredLines := make([]string, len(lines))
		for i, line := range lines {
			centeredLines[i] = centerText(line, width)
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
