package shell

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"essh/data"
)

type model struct {
	choices  []data.Target    // 可供选择的东西
	cursor   int              // 指针
	selected map[int]struct{} // 哪些items被选择了
}

var (
	page      int
	springred = lipgloss.NewStyle().Foreground(lipgloss.Color("#f1939c"))
	gold      = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
	blue      = lipgloss.NewStyle().Foreground(lipgloss.Color("#5698c3"))
)

func initialModel() model {
	return model{
		// items list
		choices: data.HostList,

		// A map which indicates which choices are selected. We're using
		// the  map like a mathematical set. The keys refer to the indexes
		// of the `choices` slice, above.
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() string {

	var builder strings.Builder

	switch page {
	// page1
	case 1:
		builder.WriteString(blue.Render("use up or down to choose you want\n\n"))

	}

	// // The header
	// s := "What should we buy at the market?\n\n"

	// // Iterate over our choices
	// for i, choice := range m.choices {

	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}

	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}

	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", gold.Render(cursor), gold.Render(checked), gold.Render(choice.User))
	// }

	// // The footer
	// s += "\nPress q to quit.\n"

	// // Send the UI for rendering
	// return s
	return builder.String()
}

func Srun() {
	page = 1
	data.InitHost()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		// fmt.Printf("Alas, there's been an error: %v", err)
		panic(fmt.Errorf("error in shellmain : %w", err))
	}
}
