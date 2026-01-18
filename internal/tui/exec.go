package tui

import (
	"os"
	"os/exec"
)

func openNvim(path string) *exec.Cmd {
	cmd := exec.Command("nvim", path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
