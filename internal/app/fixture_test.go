package app

import "github.com/umekikazuya/gh-notify/internal/notification"

var (
	repo1             = "example/ui"
	repo2             = "example/app"
	hostName          = "http://example.com"
	notificationData1 = notification.Notification{
		ID:     "1",
		Reason: "review",
		URL:    hostName + "/" + repo1,
		Number: "#12",
		Title:  "Rotate TLS certs",
		Repo:   "example/ui",
		Age:    "30m",
	}
	notificationData2 = notification.Notification{
		ID:     "2",
		Reason: "mention",
		URL:    hostName + "/" + repo2,
		Number: "#111",
		Title:  "キャッシュの\n確認\rをお願いします。",
		Repo:   "example/app",
		Age:    "1h",
	}
	notificationData3 = notification.Notification{
		ID:     "3",
		Reason: "mention",
		URL:    hostName + "/" + repo2,
		Number: "#111",
		Title:  "APIの\n確認\rをお願いします。",
		Repo:   "example/app",
		Age:    "1h",
	}
)
