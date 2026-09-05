package app

import (
	"testing"

	tea "charm.land/bubbletea/v2"
)

func Test_model_Update(t *testing.T) {
	tests := []struct {
		name     string
		seedFn   func(t *testing.T) tea.Model
		msg      tea.Msg
		assertFn func(t *testing.T, m tea.Model)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := tt.seedFn(t)
			_, _ = m.Update(tt.msg)
			tt.assertFn(t, m)
		})
	}
}
