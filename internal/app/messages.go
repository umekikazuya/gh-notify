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
		id string
	}
	MarkSuccessedMsg struct {
		Notification notification.Notification
	}
	MarkFailedMsg struct {
		Err error
	}
	OpenBrowserIdleMsg struct {
		URL string
	}
	OpenBrowserFailedMsg struct {
		Err error
	}
)
