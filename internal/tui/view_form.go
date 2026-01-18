package tui

import "github.com/charmbracelet/lipgloss"

func viewForm(m Model) string {
	title := lipgloss.NewStyle().
		Bold(true).
		Render("ADD / EDIT PROJECT")

	s := title + "\n\n"

	labels := []string{"Name", "Path", "Type"}

	views := []string{
		m.form.name.View(),
		m.form.path.View(),
		m.form.typ.View(),
	}

	for i := 0; i < 3; i++ {
		label := labels[i]

		// highlight label aktif
		if m.form.focus == i {
			label = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#7D56F4")).
				Render("> " + label)
		}

		s += label + ":\n" + views[i] + "\n\n"
	}

	if m.form.error != "" {
		s += lipgloss.NewStyle().
			Foreground(lipgloss.Color("#ff5555")).
			Render("Error: "+m.form.error) + "\n\n"
	}

	s += "[tab] next • [shift+tab] prev • [enter] save • [esc] cancel"

	return s
}
