package cmd

import (
	"dg/style"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var vcChoices = []string{"Create A New Branch", "Select An Existing Branch", "Exit"}

type vcModel struct {
	cursor         int
	choice         string
	branchName     string
	creatingBranch bool
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
			if m.creatingBranch {
				m.branchName = strings.TrimSpace(m.branchName)
				if m.branchName != "" {
					return m, tea.Quit
				}
			} else {
				m.choice = vcChoices[m.cursor]
				if m.choice == "Create A New Branch" {
					m.creatingBranch = true
					m.branchName = ""
				} else {
					return m, tea.Quit
				}
			}

		case "down", "j":
			if !m.creatingBranch {
				m.cursor++
				if m.cursor >= len(vcChoices) {
					m.cursor = 0
				}
			}

		case "up", "k":
			if !m.creatingBranch {
				m.cursor--
				if m.cursor < 0 {
					m.cursor = len(vcChoices) - 1
				}
			}

		case "backspace", "ctrl+h":
			if m.creatingBranch && len(m.branchName) > 0 {
				m.branchName = m.branchName[:len(m.branchName)-1]
			}

		default:
			if m.creatingBranch {
				m.branchName += msg.String()
			}
		}
	}

	return m, nil
}

func (m vcModel) View() string {
	s := strings.Builder{}
	s.WriteString("\n")
	if m.creatingBranch {
		s.WriteString(style.Orange.Render("Enter the new branch name: "))
		s.WriteString(m.branchName)
	} else {
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
	}
	return s.String()
}
