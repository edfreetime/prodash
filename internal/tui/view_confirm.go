package tui

func viewConfirm(m Model) string {
	p := m.projects[m.cursor]

	s := "Delete project: " + p.Name + " ?\n\n"
	s += "y = yes • n = no"

	return s
}
