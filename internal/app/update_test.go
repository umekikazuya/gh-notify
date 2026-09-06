package app

import (
	"reflect"
	"testing"

	tea "charm.land/bubbletea/v2"
	"github.com/umekikazuya/gh-notify/internal/notification"
)

func Test_model_Update(t *testing.T) {
	tests := []struct {
		name     string
		seedFn   func(t *testing.T) tea.Model
		msg      tea.Msg
		want     string
		assertFn func(t *testing.T, cmd tea.Cmd)
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
			want: ` gh notify                                                              3 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
  mention #111 APIの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
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
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
		},
		{
			name: "q キーはなにもしない",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{Text: "q", Mod: 0, Code: 'q', ShiftedCode: 0, BaseCode: 0, IsRepeat: false},
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
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
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
> review  #12  Rotate TLS certs 30m
  mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
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
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
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
			want: ` gh notify                                                              2 unread
--------------------------------------------------------------------------------
  review  #12  Rotate TLS certs 30m
> mention #111 キャッシュの 確認 をお願いします。  1h
--------------------------------------------------------------------------------
 w:open r:read u:unread q:quit`,
			assertFn: func(t *testing.T, cmd tea.Cmd) {
				t.Helper()
				got := cmd()
				if reflect.TypeOf(got) != reflect.TypeOf(tea.QuitMsg{}) {
					t.Errorf(
						"check type: got = %v, tea.QuitMsg{} = %v",
						reflect.TypeOf(got),
						reflect.TypeOf(tea.QuitMsg{}),
					)
				}
			},
		},
		{
			name: "読み込み中メッセージ",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{
					notificationData1,
					notificationData2,
				}, Cursor: 1, Width: 80, All: false, Loading: false, Error: nil}
			},
			msg: tea.KeyPressMsg{
				Text:        "ctrl+r",
				Mod:         0,
				ShiftedCode: 0,
				BaseCode:    0,
				IsRepeat:    false,
			},
			want: "",
		},
		{
			name: "読み込みの排他制御",
			seedFn: func(t *testing.T) tea.Model {
				t.Helper()
				return &Model{Notifications: []notification.Notification{notificationData1, notificationData2}, Cursor: 0, Width: 80, All: false, Loading: true, Error: nil}
			},
			msg: tea.KeyPressMsg{
				Text:        "ctrl+r",
				Mod:         0,
				ShiftedCode: 0,
				BaseCode:    0,
				IsRepeat:    false,
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.seedFn(t)
			_, cmd := m.Update(tt.msg)
			if tt.assertFn != nil {
				tt.assertFn(t, cmd)
			}
			if cmd != nil {
				cmd()
			}
			if m.View().Content != tt.want {
				t.Errorf("got = %v, want = %v", m.View().Content, tt.want)
			}
		})
	}
}
