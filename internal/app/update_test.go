package app

import (
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func Test_model_Update(t *testing.T) {
	tests := []struct {
		name          string
		seedFn        func(t *testing.T) tea.Model
		msg           tea.Msg
		wantBeforeCMD string
		wantAfterCMD  string
	}{
		{
			name: "keydown",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
					notificationData3,
				}, Cursor: 0, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "j", Mod: 0, Code: 'j', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			wantBeforeCMD: ` gh notify                                                              3 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
  mention #111 APIの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
			wantAfterCMD: ` gh notify                                                              3 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
  mention #111 APIの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
		},
		{
			name: "keyup",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "k", Mod: 0, Code: 'k', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			wantBeforeCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
			wantAfterCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
		},
		{
			name: "カーソルがこれ以上上がらない",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 0, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "k", Mod: 0, Code: 'k', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			wantBeforeCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
			wantAfterCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
		},
		{
			name: "カーソルがこれ以上下がらない",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "j", Mod: 0, Code: 'j', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			wantBeforeCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
			wantAfterCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
		},
		{
			name: "q キー入力",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "q", Mod: 0, Code: 'q', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			wantBeforeCMD: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read/unread q:quit`,
			wantAfterCMD: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.seedFn(t)
			_, cmd := m.Update(tt.msg)
			if m.View().Content != tt.wantBeforeCMD {
				t.Errorf("got = %v, wantBeforeCMD = %v", m.View().Content, tt.wantBeforeCMD)
			}
			if cmd != nil {
				cmd()
			}
			if m.View().Content != tt.wantAfterCMD {
				t.Errorf("got = %v, wantAfterCMD = %v", m.View().Content, tt.wantAfterCMD)
			}
		})
	}
}
