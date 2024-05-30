//go:build !containers && !disable_notification

package app

import notify "github.com/karapetianash/notifier"

func send_notification(msg string) {
	n := notify.New("Pomodoro", msg, notify.SeverityNormal)

	n.Send()
}
