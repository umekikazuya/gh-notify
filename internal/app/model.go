package app

import (
	"github.com/umekikazuya/gh-notify/internal/notification"
)

type Model struct {
	Notifications []notification.Notification
	Cursor        int
	Width         int
	All           bool
	Loading       bool
	Error         error
}

func NewModel(
	width int,
	all bool,
) Model {
	return Model{
		Width: width,
		All:   all,
	}
}
