package app

import "github.com/umekikazuya/gh-notify/internal/notification"

type (
	loadIdleMsg      struct{}
	loadSuccessedMsg struct {
		notifications []notification.Notification
	}
	loadFailed struct {
		Err error
	}
)
