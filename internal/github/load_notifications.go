package github

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os/exec"
	"time"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/app"
)

func FindAll() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

var _ app.FindAllFn = FindAll

func execGhApiNotifications() error {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	cmd := exec.CommandContext(ctx, "gh", "api", "notifications?all=true")
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr
	err := cmd.Run()
	if err != nil {
		return err
	}
	var data []demo
	json.Unmarshal(stdOut.Bytes(), &data)
	for _, item := range data {
		log.Printf("item = %v", item.ID)
	}

	return nil
}

type demo struct {
	ID string `json:"id"`
}
