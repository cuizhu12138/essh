package shell

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"essh/config"
)

// page1 home
type home struct {
	choices  []string         // 可供选择的东西
	cursor   int              // 指针
}

func inithome() home {
	return home{
		// items list
		choices:  []string{"add new host", "connect to host"},
	}
}

func (m home) Init() tea.Cmd {
	return nil
}

func (m home) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		// The "enter" key and the spacebar (a literal space) toggle
		// home界面选择
		case "enter", " ":
			s.page = m.cursor + 2
		}
	}
	if s.page == 2 {
		return s.Connectlist.Update(msg)
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m home) View() string {

	var builder strings.Builder

	var debug string = ""

	// The header
	builder.WriteString("use up or down to choose you want")
	builder.WriteString("\n\n")
	for i, choice := range m.choices {
		// Is the cursor pointing at this choice?
		cursor := "  " // no cursor
		if m.cursor == i {
			cursor = ">>" // cursor!
		}

		// Is this choice selected?
		// if _, ok := m.selected[i]; ok {
		// 	switch i {
		// 	case 0:
		// 		debug += "0"
		// 	case 1:
		// 		debug += "1"
		// 	}
		// }

		// Render the row
		builder.WriteString(fmt.Sprintf("%s %s\n", green.Render(cursor), gold.Render(choice)))
	}
	// debug
	if config.DebugMode {
		builder.WriteString(debug)
	}

	// The footer
	builder.WriteString("\n")
	builder.WriteString("Press q to quit.")
	builder.WriteString("\n")

	return builder.String()
}
