package tui

import (
	"os"

	"prodash/internal/config"

	tea "github.com/charmbracelet/bubbletea"
)

func updateForm(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "esc":
			m.mode = modeList
			return m, nil

		case "tab", "down":
			m.form.focus = (m.form.focus + 1) % 3

		case "shift+tab", "up":
			m.form.focus--
			if m.form.focus < 0 {
				m.form.focus = 2
			}

		case "ctrl+p":
			m.mode = modePathPicker

			// mulai dari home atau dari isi field
			start := m.form.path.Value()
			if start == "" {
				home, _ := os.UserHomeDir()
				start = home
			}

			m.picker = newPicker(start)
			return m, nil

		case "enter":
			// kalau masih di field pertama → pindah ke field berikut
			if m.form.focus < 2 {
				m.form.focus++
				return m, nil
			}

			if errMsg := validateForm(m.form); errMsg != "" {
				m.form.error = errMsg
				return m, nil
			}

			// SAVE DATA
			newP := config.Project{
				Name: m.form.name.Value(),
				Path: m.form.path.Value(),
				Type: m.form.typ.Value(),
				Commands: map[string]string{
					"open": "nvim .",
				},
			}

			if m.selected >= 0 && m.selected < len(m.projects) {
				m.projects[m.selected] = newP
			} else {
				m.projects = append(m.projects, newP)
			}

			m.cfg.Projects = m.projects
			config.Save(m.cfg)

			m.mode = modeList
			m.selected = -1

			return m, nil
		}
	}

	// kirim input ke field yg lagi fokus
	switch m.form.focus {
	case 0:
		m.form.name, cmd = m.form.name.Update(msg)
	case 1:
		m.form.path, cmd = m.form.path.Update(msg)
	case 2:
		m.form.typ, cmd = m.form.typ.Update(msg)
	}

	m.syncFocus()
	m.form.error = ""

	return m, cmd
}
