package app

import (
	"errors"
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
				Cursor:  0,
				Width:   80,
				All:     false,
				Loading: false,
				Error:   nil,
			},
			want: "",
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
			want: "Loading notifications.",
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
			want: "No unread notifications.",
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
		{
			name: "all(`--all`)",
			model: Model{
				Notifications: []notification.Notification{
					notificationData1, notificationData2,
				},
				Cursor:  0,
				Width:   80,
				All:     true,
				Loading: false,
				Error:   nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := View(tt.model)
			if got != tt.want {
				t.Errorf("View() =\n %v, want =\n %v", got, tt.want)
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
			in:   "ab  cde",
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

func Test_truncate(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "",
			in:   "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := truncate(tt.in)
			if got != tt.want {
				t.Errorf("truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}
