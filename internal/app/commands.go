package app

import (
	tea "charm.land/bubbletea/v2"
)

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

// 読み込み結果を
func loadNotifications(m *Model, client githubClient) tea.Cmd {
	return func() tea.Msg {
		ns, err := client.Exec()
		if err != nil {
			return loadFailed{
				Err: err,
			}
		}
		return loadSuccessedMsg{
			notifications: ns,
		}
	}
}
