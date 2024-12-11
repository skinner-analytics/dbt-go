/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/TheZoraiz/ascii-image-converter/aic_package"
	"github.com/charmbracelet/lipgloss"
)

// https://github.com/charmbracelet/lipgloss/blob/master/README.md

var dbt = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("32"))

// Info
var cyan = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("36"))

// Warning
var yellow = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("33"))

// Error
var red = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("31"))

// Primary
var blue = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("34"))

// Secondary
var gray = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#6C757D"))

// Light
var light_gray = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#F8F9FA"))

// Dark
var dark_gray = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#343A40"))

// dg colors
var orange = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFA500"))

var orange_indent = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FFA500")).
	PaddingTop(1).
	PaddingLeft(2).
	Width(120)

// png to ascii
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
	flags.Dimensions = []int{26, 13}      // Adjust as needed
	flags.FontColor = [3]int{255, 165, 0} // RGB for orange

	asciiArt, err := aic_package.Convert(filePath, flags)
	if err != nil {
		return fmt.Sprintf("Error generating ASCII art: %v", err)
	}
	return asciiArt
}
