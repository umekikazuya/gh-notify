package app

import (
	tea "charm.land/bubbletea/v2"
)

// Update implements [tea.Model].
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		if msg.Text == "j" {
			return m, cursorDown(m)
		}
		if msg.Text == "k" {
			return m, cursorUp(m)
		}
	}
	return m, nil
}

func cursorDown(m *Model) tea.Cmd {
	return func() tea.Msg {
		if len(m.Notifications) == m.Cursor {
			return nil
		}
		m.Cursor++
		return nil
	}
}

func cursorUp(m *Model) tea.Cmd {
	return func() tea.Msg {
		if m.Cursor == 0 {
			return nil
		}
		m.Cursor--
		return nil
	}
}
