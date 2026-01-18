package tui

import "github.com/charmbracelet/lipgloss"

func viewList(m Model) string {
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Render("ProDash - Project Dashboard")

	s := header + "\n\n"

	for i, project := range m.projects {
		cursor := "  "
		name := project.Name

		if m.cursor == i {
			cursor = "> "
			name = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#04B575")).
				Render(name)
		}

		s += cursor + name + "\n"
	}

	s += "\n[a] add • [e] edit • [d] delete • [q] quit"
	return s
}
