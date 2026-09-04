package app

import "github.com/umekikazuya/gh-notify/internal/notification"

type Model struct {
	Notifications []notification.Notification
	Cursor        int
	Width         int
	All           bool
	Loading       bool
	Error         error
}

func View(model Model) string {
	return ""
}
