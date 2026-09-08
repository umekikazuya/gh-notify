package github

import (
	"bytes"
	"context"
	"encoding/json"
	"os/exec"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/app"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func FindAll() tea.Cmd {
	return func() tea.Msg {
		ns, err := execGhApiNotifications()
		if err != nil {
			return app.LoadFailedMsg{
				Err: err,
			}
		}
		return app.LoadSuccessedMsg{
			Notifications: ns,
		}
	}
}

var _ app.FindAllFn = FindAll

func execGhApiNotifications() ([]notification.Notification, error) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(
		ctx,
		"gh",
		"api",
		"notifications?all=false",
		"--paginate",
		"--slurp",
	)
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	var page [][]thread
	err = json.Unmarshal(stdOut.Bytes(), &page)
	if err != nil {
		return nil, err
	}
	var data []thread
	for _, item := range page {
		data = append(data, item...)
	}
	ns := make([]notification.Notification, 0, len(data))
	for _, item := range data {
		n := mapNotification(item)
		ns = append(ns, n)
	}
	return ns, nil
}

type thread struct {
	ID        string        `json:"id"`
	Reason    string        `json:"reason"`
	URL       string        `json:"url"`
	UpdatedAt time.Time     `json:"updated_at"`
	Repo      threadRepo    `json:"repository"`
	Subject   threadSubject `json:"subject"`
}

type threadRepo struct {
	Name string `json:"full_name"`
}

type threadSubject struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}
