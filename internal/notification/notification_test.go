package notification

import "testing"

func TestSendNotification(t *testing.T) {
	SendNotification(Config{Title: "Test", Message: "Lorem"}, "../../assets/reminder.png")
}
