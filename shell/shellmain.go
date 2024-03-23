package shell

import (
	"essh/data"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type state struct {
	page        int
	Home        home
	Connectlist connectlist
}

var (
	//red   = lipgloss.NewStyle().Foreground(lipgloss.Color("#f1939c"))
	gold    = lipgloss.NewStyle().Foreground(lipgloss.Color("#B8860B"))
	blue    = lipgloss.NewStyle().Foreground(lipgloss.Color("#5698c3")).Width(15)
	green   = lipgloss.NewStyle().Foreground(lipgloss.Color("#45b787"))
	purple  = lipgloss.NewStyle().Foreground(lipgloss.Color("#951c48")).Width(15)
	purple2 = lipgloss.NewStyle().Foreground(lipgloss.Color("#813c85"))

	s state
)

func initstate() state {
	return state{
		page:        1,
		Home:        inithome(),
		Connectlist: initconnectlist(),
	}
}

func (m state) Init() tea.Cmd { return nil }

func (m state) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if m.page == 1 {
		return m.Home.Update(msg)
	} else if m.page == 2 {
		return m.Connectlist.Update(msg)
	}
	return m, nil
}

func (m state) View() string {
	var s string
	if m.page == 1 {
		s = m.Home.View()
	} else if m.page == 2 {
		s = m.Connectlist.View()
	}
	return s
}

func Srun() {
	data.InitHost()
	s = initstate()
	p := tea.NewProgram(&s)
	if _, err := p.Run(); err != nil {
		// fmt.Printf("Alas, there's been an error: %v", err)
		panic(fmt.Errorf("error in shellmain : %w", err))
	}
}
