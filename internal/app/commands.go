package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

type MarkType int

const (
	MarkTypeRead   MarkType = iota
	MarkTypeUnread MarkType = iota
	MarkTypeDone   MarkType = iota
)

type (
	FindAllFn              func() tea.Cmd
	MarkReadNotificationFn func(
		n notification.Notification,
	) tea.Cmd
)
