package cmd

import (
	"dg/style"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var vcChoices = []string{"Create A New Branch", "Select An Existing Branch", "Exit"}

type vcModel struct {
	cursor int
	choice string
}

func (m vcModel) Init() tea.Cmd {
	return nil
}

func (m vcModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = vcChoices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(vcChoices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(vcChoices) - 1
			}
		}
	}

	return m, nil
}

func (m vcModel) View() string {
	s := strings.Builder{}
	s.WriteString("\n")
	s.WriteString(style.Orange.Render("What would you like to do?"))
	s.WriteString("\n\n")

	for i := 0; i < len(vcChoices); i++ {
		if m.cursor == i {
			s.WriteString("(" + style.Orange.Render("*") + ") ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(vcChoices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n")
	s.WriteString(style.Orange.Render("(press q to quit)"))
	s.WriteString("\n")
	return s.String()
}
