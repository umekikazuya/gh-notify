package app

import (
	"fmt"
	"strconv"
	"strings"

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

func View(model Model) string {
	if model.Loading {
		return "Loading notifications."
	}
	if model.Error != nil {
		return fmt.Sprintf("Error: %v", model.Error.Error())
	}
	if len(model.Notifications) == 0 {
		return "No unread notifications."
	}
	lines := []string{
		" gh notify " + strconv.Itoa(len(model.Notifications)) + " unread",
		strings.Repeat("-", model.Width),
	}
	for i, n := range model.Notifications {
		var prefix string
		if i == model.Cursor {
			prefix = "> "
		} else {
			prefix = "  "
		}
		lines = append(
			lines,
			strings.Join(
				[]string{
					prefix,
					normalizeText(n.Reason),
					n.Number,
					normalizeText(n.Title),
					normalizeText(n.Age),
				},
				" ",
			),
		)
	}
	lines = append(lines, strings.Repeat("-", model.Width))
	lines = append(lines, " w:open r:read u:unread q:quit")
	return strings.Join(lines, "\n")
}

// normalizeText は以下の要素を1つの空白に正規化する
//
// - `\r`
// - `\n`
// - tab
// - 連続空白
func normalizeText(in string) string {
	return strings.Join(
		strings.Fields(in),
		" ",
	)
}

func truncate(in string) string {
	return ""
}
