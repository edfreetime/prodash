package tui

import (
	"os"

	"prodash/internal/config"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type mode int

const (
	modeList mode = iota
	modeForm
	modeConfirmDelete
	modePathPicker
)

type Model struct {
	cursor   int
	projects []config.Project

	// states
	mode     mode
	form     formModel
	picker   pickerModel
	selected int

	cfg config.Config
}

type formModel struct {
	name textinput.Model
	path textinput.Model
	typ  textinput.Model

	focus int
	error string
}

type pickerModel struct {
	currentPath string
	entries     []os.DirEntry
	cursor      int
}

func InitialModel(cfg config.Config) Model {
	return Model{
		projects: cfg.Projects,
		cfg:      cfg,
		mode:     modeList,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.mode {

	// =======================
	// MODE LIST
	// =======================
	case modeList:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {

			case "q", "ctrl+c":
				return m, tea.Quit

			case "j", "down":
				if m.cursor < len(m.projects)-1 {
					m.cursor++
				}

			case "k", "up":
				if m.cursor > 0 {
					m.cursor--
				}

			// ---- ADD ----
			case "a":
				m.mode = modeForm
				m.form = formModel{
					name:  newTextInput("Project name"),
					path:  newTextInput("Project path"),
					typ:   newTextInput("Type (go/tauri/etc)"),
					focus: 0,
				}
				m.form.name.Focus()
				return m, nil

			// ---- EDIT ----
			case "e":
				p := m.projects[m.cursor]

				m.form = formModel{
					name:  newTextInput("Project name"),
					path:  newTextInput("Project path"),
					typ:   newTextInput("Type"),
					focus: 0,
				}
				m.form.name.SetValue(p.Name)
				m.form.path.SetValue(p.Path)
				m.form.typ.SetValue(p.Type)

				m.mode = modeForm
				m.selected = m.cursor

				return m, nil

			// ---- DELETE ----
			case "d":
				if len(m.projects) == 0 {
					return m, nil
				}

				m.mode = modeConfirmDelete
				return m, nil

			case "enter":
				if len(m.projects) == 0 {
					return m, nil
				}

				p := m.projects[m.cursor]

				return m, tea.ExecProcess(
					openNvim(p.Path),
					func(err error) tea.Msg {
						return nil
					},
				)
			}
		}

	// =======================
	// MODE FORM
	// =======================
	case modeForm:
		return updateForm(m, msg)

	// =======================
	// MODE PATH PICKER
	// =======================
	case modePathPicker:
		return updatePicker(m, msg)

	// =======================
	// MODE CONFIRM DELETE
	// =======================
	case modeConfirmDelete:
		if key, ok := msg.(tea.KeyMsg); ok {
			switch key.String() {
			case "y":
				m.projects = append(
					m.projects[:m.cursor],
					m.projects[m.cursor+1:]...,
				)

				m.cfg.Projects = m.projects
				config.Save(m.cfg)

				m.mode = modeList
				if m.cursor > 0 {
					m.cursor--
				}

			case "n", "esc":
				m.mode = modeList
			}
		}
	}

	return m, nil
}

func (m *Model) syncFocus() {
	inputs := []*textinput.Model{
		&m.form.name,
		&m.form.path,
		&m.form.typ,
	}

	for i, input := range inputs {
		if i == m.form.focus {
			input.Focus()
			input.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
			input.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("15"))
		} else {
			input.Blur()
		}
	}
}

func (m Model) View() string {
	switch m.mode {

	case modeList:
		return viewList(m)

	case modeForm:
		return viewForm(m)

	case modeConfirmDelete:
		return viewConfirm(m)

	case modePathPicker:
		return viewPicker(m)

	default:
		return "unknown mode"
	}
}
