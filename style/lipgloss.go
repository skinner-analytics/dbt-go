package style

import (
	"github.com/charmbracelet/lipgloss"
)

// https://github.com/charmbracelet/lipgloss/blob/master/README.md
// https://www.rapidtables.com/web/color/

var Cyan = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FFFF"))

var Yellow = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFF00"))

var Red = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FF0000"))

var Blue = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#0000FF"))

var Green = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#00FF00"))

var Purple = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#9370DB"))

var LightGray = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#D3D3D3"))

var Gray = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#A9A9A9"))

var DarkGray = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#808080"))

// dg colors
var Dg = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#c82896"))

var DgIndent = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#c82896")).
	PaddingTop(1).
	PaddingLeft(2).
	Width(120)
