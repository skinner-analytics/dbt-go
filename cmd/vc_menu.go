package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type vcModel struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
}

func initialVcModel() vcModel {
	return vcModel{
		choices:  []string{"Create A New Branch", "Select An Existing Branch", "Exit"},
		selected: make(map[int]struct{}),
	}
}

func (m vcModel) Init() tea.Cmd {
	return nil
}

func (m vcModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			fmt.Printf("Selected option: %s\n", m.choices[m.cursor])
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m vcModel) View() string {
	s := "What would you like to do?\n\n"

	for i, choice := range m.choices {
		cursor := " " // no cursor
		if m.cursor == i {
			cursor = ">" // cursor!
		}
		checked := " " // not selected
		if _, ok := m.selected[i]; ok {
			checked = "x" // selected!
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPress q to quit.\n"

	return s
}
