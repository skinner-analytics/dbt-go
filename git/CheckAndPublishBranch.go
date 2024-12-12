package git

import (
	"dg/style"
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choice string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "y", "Y":
			m.choice = "y"
			return m, tea.Quit
		case "n", "N":
			m.choice = "n"
			return m, tea.Quit
		}
	case tea.Msg:
		return m, tea.Quit
	}
	return m, nil
}

func (m model) View() string {
	return style.Gray.Render("Branch is not published. Would you like to publish it to the remote? (y/n)")
}

func CheckAndPublishBranch() error {
	currentBranch, err := GetCurrentBranch()
	if err != nil {
		p := tea.NewProgram(model{})
		m, err := p.Run()
		if err != nil {
			return fmt.Errorf(style.Red.Render("Prompt failed: %v"), err)
		}

		if m.(model).choice == "y" {
			pushCmd := exec.Command("git", "push", "--set-upstream", "origin", currentBranch)
			pushOutput, err := pushCmd.CombinedOutput()
			if err != nil {
				return fmt.Errorf("error pushing changes: %v\n%s", err, pushOutput)
			}
			fmt.Println("Successfully pushed the changes to the remote.")
		}
	} else {
		pushCmd := exec.Command("git", "push")
		pushOutput, err := pushCmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("error pushing changes: %v\n%s", err, pushOutput)
		}
		fmt.Println("Successfully pushed the changes to the remote.")
	}
	return nil
}
