package app

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func Test_model_Update(t *testing.T) {
	tests := []struct {
		name     string
		seedFn   func(t *testing.T) tea.Model
		msg      tea.Msg
		assertFn func(t *testing.T, m tea.Model)
	}{
		{
			name: "",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{notificationData1, notificationData2}, Cursor: 0, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg:      tea.KeyPressMsg{Text: "j", Mod: 0, Code: 'j', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			assertFn: func(t *testing.T, m tea.Model) { t.Helper(); t.Logf("content\n %v", m.View().Content); t.Fatal("") },
		},
		{
			name: "",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "k", Mod: 0, Code: 'k', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			assertFn: func(t *testing.T, m tea.Model) {
				t.Helper()
				t.Logf("content\n %v", m.View().Content)
				t.Fatal("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.seedFn(t)
			_, cmd := m.Update(tt.msg)
			if cmd != nil {
				cmd()
			}
			tt.assertFn(t, m)
		})
	}
}
