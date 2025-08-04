package main

import (
	"fmt"
	"opsdash/db"
	"opsdash/tui"
	"os"

	"github.com/charmbracelet/bubbletea"
)

func main() {
	db.InitDB("opsdash.db")

	p := tea.NewProgram(tui.InitialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting TUI: %v\n", err)
		os.Exit(1)
	}
}
