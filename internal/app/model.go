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
	FindAllFn     FindAllFn
}

type (
	FindAllFn            func() tea.Cmd
	ReadNotificationFn   func(notification.Notification) tea.Cmd
	UnreadNotificationFn func(notification.Notification) tea.Cmd
)

func NewModel(
	width int,
	all bool,
	findAllFn FindAllFn,
) Model {
	return Model{
		Width:     width,
		All:       all,
		FindAllFn: findAllFn,
	}
}

// Init implements [tea.Model].
func (m *Model) Init() tea.Cmd {
	return func() tea.Msg {
		return m.FindAllFn
	}
}

var _ tea.Model = (*Model)(nil)
