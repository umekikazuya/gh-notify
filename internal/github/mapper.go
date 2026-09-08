package github

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/umekikazuya/gh-notify/internal/notification"
)

func mapNotification(raw thread) notification.Notification {
	return notification.Notification{
		ID:     raw.ID,
		Reason: raw.Reason,
		URL:    raw.Subject.URL,
		Number: formatNumber(raw.Subject.URL),
		Title:  raw.Subject.Title,
		Repo:   raw.Repo.Name,
		Age:    formatAge(raw.UpdatedAt, func() time.Time { return time.Now() }),
	}
}

func formatNumber(in string) string {
	parsed, err := url.Parse(in)
	if err != nil {
		return ""
	}
	parts := strings.Split(strings.Trim(parsed.Path, "/"), "/")
	return "#" + parts[len(parts)-1]
}

// formatAge は time.Time 型を以下形式に変換する
//
// 1分~59分前: %dm age
// 1時間~24時間前: %dh age
// 24時間以上: %dd age
func formatAge(in time.Time, baseTimeFn func() time.Time) string {
	diff := baseTimeFn().Sub(in)
	if diff < 0 {
		return ""
	}
	switch {
	case diff < time.Hour:
		return fmt.Sprintf("%dm age", diff/time.Minute)
	case diff < 24*time.Hour:
		return fmt.Sprintf("%dh age", diff/time.Hour)
	default:
		return fmt.Sprintf("%dd age", diff/(24*time.Hour))
	}
}
