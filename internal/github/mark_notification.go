package github

import (
	"bytes"
	"context"
	"os/exec"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/app"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func MarkReadNotification(
	n notification.Notification,
) tea.Cmd {
	return func() tea.Msg {
		err := execGhApiThread(n.ID)
		if err != nil {
			return app.MarkFailedMsg{
				Err: err,
			}
		}
		return app.MarkSuccessedMsg{
			Notification: notification.Notification{
				ID:     n.ID,
				Reason: n.Reason,
				URL:    n.URL,
				Number: n.Number,
				Title:  n.Title,
				Repo:   n.Repo,
				Age:    n.Age,
			},
		}
	}
}

var _ app.MarkReadNotificationFn = MarkReadNotification

func execGhApiThread(threadID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// #nosec G204
	cmd := exec.CommandContext(
		ctx,
		"gh",
		"api",
		"--method", "PATCH",
		"/notifications/threads/"+threadID,
	)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
