package app

import (
	tea "charm.land/bubbletea/v2"
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

// Init implements [tea.Model].
func (m *Model) Init() tea.Cmd {
	return func() tea.Msg {
		return loadNotifications(nil)
	}
}

var _ tea.Model = (*Model)(nil)
