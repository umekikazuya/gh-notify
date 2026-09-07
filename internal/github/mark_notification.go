package github

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/app"
)

func MarkNotification(threadID string,
	markType app.MarkType,
) tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

var _ app.MarkNotificationFn = MarkNotification
