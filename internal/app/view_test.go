package app

import (
	"errors"
	"strings"
	"testing"

	"github.com/umekikazuya/gh-notify/internal/notification"
)

func TestView(t *testing.T) {
	tests := []struct {
		name  string
		model Model
		want  string
	}{
		{
			name: "default",
			model: Model{
				Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				},
				Cursor:  1,
				Width:   80,
				All:     false,
				Loading: false,
				Error:   nil,
			},
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12   Rotate TLS certs 30m
> mention #111  キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 o:open m:markRead/markUnread q:quit`,
		},
		{
			name: "loading",
			model: Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         80,
				All:           false,
				Loading:       true,
				Error:         nil,
			},
			want: ` gh notify                                                              0 unread
--------------------------------------------------------------------------------
Loading notifications.
--------------------------------------------------------------------------------
 o:open m:markRead/markUnread q:quit`,
		},
		{
			name: "notfound",
			model: Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         80,
				All:           false,
				Loading:       false,
				Error:         nil,
			},
			want: ` gh notify                                                              0 unread
--------------------------------------------------------------------------------
No notifications.
--------------------------------------------------------------------------------
 o:open m:markRead/markUnread q:quit`,
		},
		{
			name: "Error(api)",
			model: Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         80,
				All:           false,
				Loading:       false,
				Error:         errors.New("APIエラー"),
			},
			want: "Error: APIエラー",
		},
		// TODO: `--all`
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := View(tt.model)
			if got != tt.want {
				t.Errorf("View() =\n%v, want =\n%v", got, tt.want)
			}
		})
	}
}

func Test_normalizeText(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "abcde",
			in:   "abcde",
			want: "abcde",
		},
		{
			name: "ab<space><space>cde",
			in:   "ab  cde",
			want: "ab cde",
		},
		{
			name: "ab<space><space><space>cde",
			in:   "ab   cde",
			want: "ab cde",
		},
		{
			name: "ab<tab>cde",
			in:   "ab	cde",
			want: "ab cde",
		},
		{
			name: "ab<\r>cde",
			in:   "ab\rcde",
			want: "ab cde",
		},
		{
			name: "ab<\n>cde",
			in:   "ab\ncde",
			want: "ab cde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := normalizeText(tt.in)
			if got != tt.want {
				t.Errorf("normalizeText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rule(t *testing.T) {
	tests := []struct {
		name  string
		width int
		want  string
	}{
		{
			name:  "80",
			width: 80,
			want:  strings.Repeat("-", 80),
		},
		{
			name:  "0",
			width: 0,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rule(tt.width)
			if got != tt.want {
				t.Errorf("rule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_displayWidth(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{
			name: "aiueo",
			in:   strings.Repeat("-", 80),
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := displayWidth(tt.in)
			if got != tt.want {
				t.Errorf("displayWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_padRight(t *testing.T) {
	tests := []struct {
		name  string
		value string
		width int
		want  string
	}{
		{
			name:  "",
			value: "abcde",
			width: 6,
			want:  "abcde ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := padRight(tt.value, tt.width)
			if got != tt.want {
				t.Errorf("padRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
