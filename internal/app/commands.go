package app

import (
	"context"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

type MarkType int

const (
	MarkTypeRead   MarkType = iota
	MarkTypeUnread MarkType = iota
	MarkTypeDone   MarkType = iota
)

type (
	FindAllFn              func() tea.Cmd
	MarkReadNotificationFn func(
		n notification.Notification,
	) tea.Cmd
	GetTreadHTMLURLFn func(rawURL string) (string, error)
)

func openBrowserCmd(
	rawURL string,
	getThreadHTMLURLFn func(rawURL string) (string, error),
) tea.Cmd {
	return func() tea.Msg {
		url, err := getThreadHTMLURLFn(rawURL)
		if err != nil {
			return err
		}
		cmd := exec.CommandContext(
			context.Background(),
			"open",
			url,
		)
		err = cmd.Start()
		if err != nil {
			return nil // TODO
		}
		return nil
	}
}
