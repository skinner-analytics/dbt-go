package cmd

import (
	"dg/style"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(infoCmd)
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show Additional Developer Information About dbt-go",
	Long: func() string {
		asciiArt := style.GetASCIIArt()
		copyright := "Copyright © 2024 Matthew Skinner"
		contact := "matthew@skinnerdev.com"

		width := 80

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
`, centeredAsciiArt, centeredCopyright, centeredContact)
	}(),
}
