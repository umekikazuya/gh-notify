package github

import (
	"strconv"
	"testing"
	"time"
)

var (
	baseTime          = time.Now()
	fixtureBaseTimeFn = func() time.Time { return baseTime }
)

func Test_formatAge(t *testing.T) {
	tests := []struct {
		in         time.Time
		baseTimeFn func() time.Time
		want       string
	}{
		{
			in:         baseTime.Add(1 * time.Minute),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "1m age",
		},
		{
			in:         baseTime.Add(59 * time.Minute),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "59m age",
		},
		{
			in:         baseTime.Add(60 * time.Minute),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "1h age",
		},
		{
			in:         baseTime.Add(61 * time.Minute),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "1h age",
		},
		{
			in:         baseTime.Add(2 * time.Hour),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "2h age",
		},
		{
			in:         baseTime.Add(24 * time.Hour),
			baseTimeFn: fixtureBaseTimeFn,
			want:       "1d age",
		},
	}
	for i, tt := range tests {
		t.Run(
			"#"+strconv.Itoa(i),
			func(t *testing.T) {
				got := formatAge(tt.in, tt.baseTimeFn)
				if got != tt.want {
					t.Errorf("formatAge() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func Test_formatNumber(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			in:   "https://api.github.com/repos/octokit/octokit.rb/issues/123",
			want: "#123",
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := formatNumber(tt.in)
			if got != tt.want {
				t.Errorf("formatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
