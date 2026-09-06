package app

import "github.com/umekikazuya/gh-notify/internal/notification"

type githubClient interface {
	FindAll() ([]Model, error)
	ReadNotification(notification.Notification) error
	UnreadNotification(notification.Notification) error
}
