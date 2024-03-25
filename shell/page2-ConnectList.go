package shell

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"essh/config"
)

// page2 connectlist
type connectlist struct {
	choices  []Target    // 可供选择的东西
	cursor   int              // 指针
	selected map[int]struct{} // 哪些items被选择了
}

func Initconnectlist() connectlist {
	return connectlist{
		// items list
		choices:  HostList,
		cursor:   0,
		selected: make(map[int]struct{}),
	}
}

func (m connectlist) Init() tea.Cmd {
	return nil
}

func (m connectlist) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.selected[m.cursor]
			if ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		case "l":
			S.page = 1
			return S, nil
		}
	}
	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m connectlist) View() string {

	var (
		builder strings.Builder
		debug   string
		which   int
	)

	// The header
	builder.WriteString("Use up or down to choose you want")
	builder.WriteString("\n\n")
	for i, choice := range m.choices {
		// Is the cursor pointing at this choice?
		cursor := "  " // no cursor
		if m.cursor == i {
			cursor = ">>" // cursor!
			which = i
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
		builder.WriteString(fmt.Sprintf("%s %s%s%s\n", green.Render(cursor), blue.Render(choice.User), purple.Render(choice.Host.Address), purple2.Render(choice.Description)))
	}

	// which 默认值
	builder.WriteString("--------- SSH Details ---------\n")
	builder.WriteString(fmt.Sprintf("%-17s%s\n%-17s%s\n%-17s%s\n%-17s%s\n", "Host", m.choices[which].Host.Address, "User", m.choices[which].User, "Port", fmt.Sprintf("%d", m.choices[which].Host.Port), "Description", m.choices[which].Description))

	// debug
	if config.DebugMode {
		builder.WriteString(debug)
	}

	// The footer
	builder.WriteString("\n")
	builder.WriteString("Press q to quit, l to last menu")
	builder.WriteString("\n")

	return builder.String()
}
