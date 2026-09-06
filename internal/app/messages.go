package app

import "github.com/umekikazuya/gh-notify/internal/notification"

type (
	loadIdleMsg      struct{}
	loadSuccessedMsg struct {
		notifications []notification.Notification
	}
	loadFailedMsg struct {
		Err error
	}
	markIdleMsg struct {
		id       string
		markType MarkType
	}
	markSuccessedMsg struct {
		n notification.Notification
	}
	markFailedMsg struct {
		Err error
	}
)
