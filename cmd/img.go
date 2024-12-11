/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func stripANSI(str string) string {
	return ansiRegex.ReplaceAllString(str, "")
}

func centerText(text string, width int) string {
	stripped := stripANSI(text)
	if len(stripped) >= width {
		return text
	}
	spaces := width - len(stripped)
	left := spaces / 2
	return strings.Repeat(" ", left) + text
}

func getASCIIArt() string {
	filePath := "img/logo.png"
	flags := aic_package.DefaultFlags()
	flags.Dimensions = []int{30, 15} // Adjust as needed
	flags.Colored = true

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		return fmt.Sprintf("Error generating ASCII art: %v", err)
	}
	return asciiArt
}
