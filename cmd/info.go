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
	Long:  "",
	RunE:  runInfo,
}

func runInfo(cmd *cobra.Command, args []string) error {
	asciiArt := style.GetASCIIArt()
	copyright := "Copyright © 2024"
	contact := "Matthew Skinner -- matthew@skinnerdev.com"
	github := "github.com/cognite-analytics/dbt-go"
	width := 80

	centeredCopyright := style.CenterText(copyright, width)
	centeredContact := style.CenterText(contact, width)
	centeredLink := style.CenterText(github, width)
	lines := strings.Split(asciiArt, "\n")
	centeredLines := make([]string, len(lines))
	for i, line := range lines {
		centeredLines[i] = style.CenterText(line, width)
	}

	centeredAsciiArt := strings.Join(centeredLines, "\n")

	fmt.Printf(`

%s

%s
%s
%s
`, centeredAsciiArt, centeredCopyright, centeredContact, centeredLink)
	return nil
}
