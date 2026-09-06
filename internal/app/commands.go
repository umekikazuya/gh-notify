package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
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

// getAllNotifications は読み込みコマンド
func getAllNotifications(client githubClient) tea.Cmd {
	return func() tea.Msg {
		ns := []notification.Notification{}
		var err error
		// ns, err := client.Exec()
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

func togglReadNotification(client githubClient) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}
