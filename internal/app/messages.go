package app

import "github.com/umekikazuya/gh-notify/internal/notification"

type (
	LoadIdleMsg      struct{}
	LoadSuccessedMsg struct {
		Notifications []notification.Notification
	}
	LoadFailedMsg struct {
		Err error
	}
	MarkIdleMsg struct {
		id       string
		markType MarkType
	}
	MarkSuccessedMsg struct {
		Notification notification.Notification
	}
	MarkFailedMsg struct {
		Err error
	}
)
