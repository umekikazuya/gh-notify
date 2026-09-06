package app

import (
	tea "charm.land/bubbletea/v2"
)

// Update implements [tea.Model].
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case loadIdleMsg:
		if m.Loading {
			return m, nil
		}
		m.Loading = true
		m.Error = nil
		return m, m.FindAllFn()
	case loadSuccessedMsg:
		m.Loading = false
		m.Error = nil
		m.Notifications = append(m.Notifications, msg.notifications...)
		return m, nil
	case loadFailed:
		m.Loading = false
		m.Error = msg.Err
		return m, nil
	case tea.KeyPressMsg:
		return handleKey(m, msg)
	default:
		return nil, nil
	}
}

func handleKey(m *Model, msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.Key().String() {
	case "m":
		if m.Loading {
			return m, nil
		}
		m.Loading = true
		return m, m.MarkNotificationFn(
			m.Notifications[m.Cursor].ID,
			markType(m.Notifications[m.Cursor]),
		)
	case "j", "down":
		if m.Cursor < len(m.Notifications)-1 {
			m.Cursor++
		}
	case "k", "up":
		if 0 < m.Cursor {
			m.Cursor--
		}
	case "q":
		return m, tea.Quit
	}
	return m, nil
}
