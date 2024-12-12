package git

import (
	"fmt"
	"os/exec"
)

func FetchChanges() error {
	fetchCmd := exec.Command("git", "fetch")
	if err := fetchCmd.Run(); err != nil {
		return fmt.Errorf("error fetching changes: %v", err)
	}
	return nil
}
