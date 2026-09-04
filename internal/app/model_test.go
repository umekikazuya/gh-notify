package app_test

import (
	"testing"

	"github.com/umekikazuya/gh-notify/internal/app"
)

func TestView(t *testing.T) {
	tests := []struct {
		name  string
		model app.Model
		want  string
	}{
		{},
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
