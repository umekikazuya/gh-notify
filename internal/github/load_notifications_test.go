package github

import "testing"

func Test_execGhApiNotifications(t *testing.T) {
	err := execGhApiNotifications()
	if err != nil {
		t.Errorf("err = %#v", err)
	}
}
