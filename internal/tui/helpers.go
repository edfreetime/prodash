package tui

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
)

func newTextInput(placeholder string) textinput.Model {
	t := textinput.New()
	t.Placeholder = placeholder
	t.Focus()
	t.CharLimit = 156
	t.Width = 40

	return t
}

func validateForm(f formModel) string {
	if f.name.Value() == "" {
		return "Name tidak boleh kosong"
	}

	if f.path.Value() == "" {
		return "Path tidak boleh kosong"
	}

	// cek folder beneran ada
	if _, err := os.Stat(f.path.Value()); os.IsNotExist(err) {
		return "directory path not found"
	}

	return ""
}

func newPicker(path string) pickerModel {
	entries, _ := os.ReadDir(path)

	return pickerModel{
		currentPath: path,
		entries:     entries,
		cursor:      0,
	}
}
