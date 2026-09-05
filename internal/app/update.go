package app

import (
	tea "charm.land/bubbletea/v2"
)

// Update implements [tea.Model].
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.Text == "j" {
			return m, moveCursorDown(m)
		}
		if msg.Text == "k" {
			return m, moveCursorUp(m)
		}
	}
	return m, nil
}

func moveCursorDown(m *Model) tea.Cmd {
	return func() tea.Msg {
		if len(m.Notifications)-1 == m.Cursor {
			return nil
		}
		m.Cursor++
		return nil
	}
}

func moveCursorUp(m *Model) tea.Cmd {
	return func() tea.Msg {
		if m.Cursor == 0 {
			return nil
		}
		m.Cursor--
		return nil
	}
}
