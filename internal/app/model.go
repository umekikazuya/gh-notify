package app

import (
	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

type Model struct {
	Notifications      []notification.Notification
	Cursor             int
	Width              int
	All                bool
	Loading            bool
	Error              error
	FindAllFn          FindAllFn
	MarkNotificationFn MarkReadNotificationFn
	GetTreadHTMLURL    GetTreadHTMLURLFn
}

func NewModel(
	width int,
	all bool,
	findAllFn FindAllFn,
	markNotificationFn MarkReadNotificationFn,
	getTreadHTMLURL GetTreadHTMLURLFn,
) tea.Model {
	return &Model{
		Width:              width,
		All:                all,
		FindAllFn:          findAllFn,
		MarkNotificationFn: markNotificationFn,
		GetTreadHTMLURL:    getTreadHTMLURL,
	}
}

// Init implements [tea.Model].
func (m *Model) Init() tea.Cmd {
	return func() tea.Msg {
		return LoadIdleMsg{}
	}
}

var _ tea.Model = (*Model)(nil)
