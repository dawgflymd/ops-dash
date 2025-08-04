package main

import (
	"fmt"
	"os"
	"github.com/charmbracelet/bubbletea"
	"opsdash/tui"
	"opsdash/db"
)

func main() {
	db.InitDB("opsdash.db")

	p := tea.NewProgram(tui.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting TUI: %v\n", err)
		os.Exit(1)
	}
}
