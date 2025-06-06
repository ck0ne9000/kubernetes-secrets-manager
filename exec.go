package main

import (
	"fmt"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

func checkSopsInstalled() error {
	cmd := exec.Command("sops", "--version")
	_, err := cmd.CombinedOutput()
	return err
}

// sops --config .sops.yaml -e --in-place secretFilePath
func encryptFileWithSops(secretName string) tea.Cmd {
	fileName := fmt.Sprintf("%s.yaml", secretName)
	return func() tea.Msg {
		cmd := exec.Command("sops", "-e", "--in-place", fileName)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}
		return out
	}
}
