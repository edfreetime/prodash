package tui

func viewPicker(m Model) string {
	s := "PICK PROJECT PATH\n"
	s += "Current: " + m.picker.currentPath + "\n\n"

	for i, e := range m.picker.entries {

		prefix := "  "
		if i == m.picker.cursor {
			prefix = "> "
		}

		name := e.Name()
		if e.IsDir() {
			name += "/"
		}

		s += prefix + name + "\n"
	}

	s += "\nenter: open • backspace: up • space: select • esc: cancel"

	return s
}
