package app_test

import (
	"errors"
	"testing"

	"github.com/umekikazuya/gh-notify/internal/app"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func TestView(t *testing.T) {
	tests := []struct {
		name  string
		model app.Model
		want  string
	}{
		{
			name: "default",
			model: app.Model{
				Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				},
				Cursor:  0,
				Width:   0,
				All:     false,
				Loading: false,
				Error:   nil,
			},
			want: "",
		},
		{
			name: "loading",
			model: app.Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         0,
				All:           false,
				Loading:       true,
				Error:         nil,
			},
			want: "Loading notifications.",
		},
		{
			name: "notfound",
			model: app.Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         0,
				All:           false,
				Loading:       false,
				Error:         nil,
			},
			want: "No unread notifications.",
		},
		{
			name: "Error(api)",
			model: app.Model{
				Notifications: []notification.Notification{},
				Cursor:        0,
				Width:         0,
				All:           false,
				Loading:       false,
				Error:         errors.New("APIエラー"),
			},
			want: "Error: APIエラー",
		},
		{
			name: "all(`--all`)",
			model: app.Model{
				Notifications: []notification.Notification{
					notificationData1, notificationData2,
				},
				Cursor:  0,
				Width:   0,
				All:     false,
				Loading: false,
				Error:   nil,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := app.View(tt.model)
			if got != tt.want {
				t.Errorf("View() = %v, want %v", got, tt.want)
			}
		})
	}
}
