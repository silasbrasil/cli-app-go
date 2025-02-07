package questions

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type EnvModel struct {
	Title    string
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
}

func (m EnvModel) Init() tea.Cmd {
	return nil
}

func (m EnvModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "enter":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case " ":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}
		}
	}

	return m, nil
}

func (m EnvModel) View() string {
	// The header
	s := m.Title + "\n\n"

	// Iterate over our choices
	for i, choice := range m.Choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.Cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.Selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// The footer
	s += "\nPressione Enter para sair.\n"

	// Send the UI for rendering
	return s
}
