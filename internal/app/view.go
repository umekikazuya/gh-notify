package app

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/umekikazuya/gh-notify/internal/notification"
)

var (
	defaultWidth = 80
	reasonWidth  = 7
	numberWidth  = 4
	ageWidth     = 3
)

func View(model Model) string {
	width := model.Width
	if width <= 0 {
		width = defaultWidth
	}
	lines := []string{
		renderHeader(width, model.All, len(model.Notifications)),
		rule(width),
	}

	switch {
	case model.Loading:
		lines = append(lines, fitLine("Loading notifications.", width))
	case len(model.Notifications) == 0:
		lines = append(lines, fitLine("No unread notifications.", width))
	default:
		for i, n := range model.Notifications {
			lines = append(
				lines,
				renderNotification(n, i == model.Cursor, width),
			)
		}
	}

	if model.Error != nil {
		return fmt.Sprintf("Error: %v", model.Error.Error())
	}
	lines = append(lines, rule(width))
	lines = append(lines, fitLine(" w:open r:read u:unread q:quit", width))
	return strings.Join(lines, "\n")
}

func renderHeader(width int, all bool, count int) string {
	command := " gh notify"
	if all {
		command = command + " --all"
	}
	countLabel := fmt.Sprintf("%d unread", count)
	space := width - displayWidth(command) - displayWidth(countLabel)
	return command + strings.Repeat(" ", space) + countLabel
}

func renderNotification(
	item notification.Notification,
	selected bool,
	width int,
) string {
	var cursor string
	if selected {
		cursor = "> "
	} else {
		cursor = "  "
	}
	reason := padRight(truncateDisplay(normalizeText(item.Reason), reasonWidth), reasonWidth)
	number := padRight(truncateDisplay(normalizeText(item.Number), numberWidth), numberWidth)
	age := padLeft(truncateDisplay(normalizeText(item.Age), ageWidth), ageWidth)

	titleWidth := width -
		displayWidth(cursor) -
		displayWidth(reason) -
		displayWidth(number) -
		displayWidth(age) - 3
	title := truncateDisplay(normalizeText(item.Title), titleWidth)

	line := cursor + reason + " " + number + " " + title + " " + age
	return fitLine(line, width)
}

// normalizeText は以下の要素を1つの空白に正規化する
//
// - `\r`
// - `\n`
// - tab
// - 連続空白
func normalizeText(in string) string {
	return strings.Join(
		strings.FieldsFunc(in, unicode.IsSpace),
		" ",
	)
}

func rule(width int) string {
	return strings.Repeat("-", max(width, 0))
}

func fitLine(value string, width int) string {
	return truncateDisplay(value, max(width, 0))
}

func truncateDisplay(value string, width int) string {
	if width <= 0 {
		return ""
	}
	if displayWidth(value) <= width {
		return value
	}
	if width <= 3 {
		return strings.Repeat(".", width)
	}
	var builder strings.Builder
	used := 0
	for _, r := range value {
		runeWidth := runeDisplayWidth(r)
		if used+runeWidth > width-3 {
			break
		}
		builder.WriteRune(r)
		used += runeWidth
	}
	return builder.String() + "..."
}

func displayWidth(in string) int {
	width := 0
	for _, r := range in {
		width += runeDisplayWidth(r)
	}
	return width
}

func runeDisplayWidth(r rune) int {
	switch {
	case r == 0:
		return 0
	case unicode.Is(unicode.Mn, r), unicode.Is(unicode.Me, r):
		return 0
	case r < utf8.RuneSelf:
		return 1
	case isWideRune(r):
		return 2
	default:
		return 1
	}
}

func isWideRune(r rune) bool {
	return r >= 0x1100 && (r <= 0x115F ||
		r == 0x2329 ||
		r == 0x232A ||
		(r >= 0x2E80 && r <= 0xA4CF && r != 0x303F) ||
		(r >= 0xAC00 && r <= 0xD7A3) ||
		(r >= 0xF900 && r <= 0xFAFF) ||
		(r >= 0xFE10 && r <= 0xFE19) ||
		(r >= 0xFE30 && r <= 0xFE6F) ||
		(r >= 0xFF00 && r <= 0xFF60) ||
		(r >= 0xFFE0 && r <= 0xFFE6) ||
		(r >= 0x20000 && r <= 0x2FFFD) ||
		(r >= 0x30000 && r <= 0x3FFFD))
}

func padRight(value string, width int) string {
	space := max(width-displayWidth(value), 0)
	return value + strings.Repeat(" ", space)
}

func padLeft(value string, width int) string {
	space := max(width-displayWidth(value), 0)
	return strings.Repeat(" ", space) + value
}
