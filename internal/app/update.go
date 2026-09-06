package app

import (
	"slices"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

// Update implements [tea.Model].
//
// TODO: 再読み込み
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
	case loadFailedMsg:
		m.Loading = false
		m.Error = msg.Err
		return m, nil
	case markIdleMsg:
		if m.Loading {
			return m, nil
		}
		m.Loading = true
		m.Error = nil
		return m, m.MarkNotificationFn(msg.id, msg.markType)
	case markSuccessedMsg:
		m.Loading = false
		m.Error = nil
		idx := slices.IndexFunc(m.Notifications, func(e notification.Notification) bool {
			return e.ID == msg.n.ID
		})
		if idx == -1 {
			return m, nil
		}
		m.Notifications[idx] = msg.n
		return m, nil
	case markFailedMsg:
		m.Loading = false
		m.Error = msg.Err
		return m, nil
	case tea.KeyPressMsg:
		return handleKey(m, msg)
	default:
		return m, nil
	}
}

func handleKey(m *Model, msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch msg.Key().String() {
	case "m":
		if len(m.Notifications) == 0 {
			return m, nil
		}
		if m.Loading {
			return m, nil
		}
		n := m.Notifications[m.Cursor]
		return m, func() tea.Msg {
			return markIdleMsg{
				id:       n.ID,
				markType: markType(n),
			}
		}
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
