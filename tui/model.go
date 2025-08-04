package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	width        int
	height       int
	commandInput textinput.Model
	commandMode  bool
	log          string
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = ":"
	ti.Focus()
	ti.CharLimit = 256
	ti.Width = 50

	return Model{
		commandInput: ti,
		commandMode:  false,
		log:          "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case ":":
			m.commandMode = true
			return m, nil

		case "enter":
			if m.commandMode {
				cmdText := strings.TrimSpace(m.commandInput.Value())

				switch cmdText {
				case "q", "q!", "wq", "x":
					return m, tea.Quit
				default:
					m.log = fmt.Sprintf("Ran command: %s", cmdText)
					m.commandInput.Reset()
					m.commandMode = false
					return m, nil
				}
			}

		case "esc":
			m.commandMode = false
			return m, nil

		case "ctrl+c":
			return m, tea.Quit
		}

		if m.commandMode {
			var cmd tea.Cmd
			m.commandInput, cmd = m.commandInput.Update(msg)
			return m, cmd
		}
	}

	return m, nil
}

func (m Model) View() string {
	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("229")).
		Bold(true).
		Padding(0, 1)

	sectionTitle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("111")).
		Bold(true).
		MarginBottom(1)

	panelStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Margin(0, 1)

	widePanel := panelStyle.Copy().
		Width(100)

	top3 := "• Task 1\n• Task 2\n• Task 3"
	messages := "• Email from Jane\n• Slack from Dev Team"
	logOut := m.log

	top3Panel := panelStyle.Render(sectionTitle.Render("📌 Top 3 Tasks") + "\n" + top3)
	logPanel := panelStyle.Render(sectionTitle.Render("🔧 Output Log") + "\n" + logOut)
	msgPanel := widePanel.Render(sectionTitle.Render("💬 Recent Messages") + "\n" + messages)

	topRow := lipgloss.JoinHorizontal(lipgloss.Top, top3Panel, logPanel)
	body := lipgloss.JoinVertical(lipgloss.Left, topRow, msgPanel)

	if m.commandMode {
		commandPrompt := sectionTitle.Render("⌨ Command:") + "\n" + m.commandInput.View()
		return lipgloss.JoinVertical(lipgloss.Left,
			titleStyle.Render("🧠 Ops Dashboard"),
			body,
			commandPrompt,
		)
	}

	return lipgloss.JoinVertical(lipgloss.Left,
		titleStyle.Render("🧠 Ops Dashboard"),
		body,
	)
}
