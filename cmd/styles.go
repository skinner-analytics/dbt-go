/*
Copyright © 2024 Matthew Skinner matthew@skinnerdev.com
*/
package cmd

import "github.com/charmbracelet/lipgloss"

// https://github.com/charmbracelet/lipgloss/blob/master/README.md

var root = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF694A"))

var dbt = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FF694A")).
	PaddingTop(1).
	PaddingLeft(2).
	Width(120)
