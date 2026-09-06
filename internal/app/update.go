package app

import (
	tea "charm.land/bubbletea/v2"
)

// Update implements [tea.Model].
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.Key().String() {
		case "r":
			if m.Loading {
				return m, nil
			}
			m.Loading = true
			return m, togglReadNotification(nil)
		case "ctrl+r":
			if m.Loading {
				return m, nil
			}
			m.Loading = true
			return m, getAllNotifications(nil)
		case "j", "down":
			if m.Cursor < len(m.Notifications)-1 {
				m.Cursor++
			}
			return m, nil
		case "k", "up":
			if 0 < m.Cursor {
				m.Cursor--
			}
			return m, nil
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}
