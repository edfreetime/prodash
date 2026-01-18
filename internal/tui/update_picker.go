package tui

import (
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
)

func updatePicker(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {

		case "esc":
			m.mode = modeForm
			return m, nil

		case "up", "k":
			if m.picker.cursor > 0 {
				m.picker.cursor--
			}

		case "down", "j":
			if m.picker.cursor < len(m.picker.entries)-1 {
				m.picker.cursor++
			}

		// masuk folder
		case "enter":
			e := m.picker.entries[m.picker.cursor]

			if e.IsDir() {
				newPath := filepath.Join(
					m.picker.currentPath,
					e.Name(),
				)

				m.picker = newPicker(newPath)
			}

		// naik folder
		case "backspace":
			parent := filepath.Dir(m.picker.currentPath)
			m.picker = newPicker(parent)

		// PILIH FOLDER INI
		case " ":
			m.form.path.SetValue(m.picker.currentPath)
			m.mode = modeForm
		}
	}

	return m, nil
}
