package main

import (
	"fmt"
	"os"

	"prodash/internal/config"
	"prodash/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("error load config:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(tui.InitialModel(cfg))

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
