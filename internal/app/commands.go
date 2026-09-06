package app

import tea "charm.land/bubbletea/v2"

type MarkType int

const (
	MarkTypeRead   MarkType = iota
	MarkTypeUnread MarkType = iota
	MarkTypeDone   MarkType = iota
)

type (
	FindAllFn          func() tea.Cmd
	MarkNotificationFn func(
		threadID string,
		markType MarkType,
	) tea.Cmd
)
