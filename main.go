package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "lsb" {
		listChangedFiles()
	}

}

func listChangedFiles() {
	cmd := exec.Command("git", "diff", "--name-only", "HEAD", "origin/main")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing git diff:", err)
		return
	}

	files := strings.Split(string(output), "\n")
	for _, file := range files {
		if strings.HasSuffix(file, ".sql") || strings.HasSuffix(file, ".yml") {
			fmt.Println(file)
		}
	}
}
