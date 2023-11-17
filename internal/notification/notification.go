package notification

import "github.com/gen2brain/beeep"

type Config struct {
	Title   string
	Message string
}

func SendNotification(config Config, iconPath ...string) {
	appIcon := "assets/reminder.png"
	if len(iconPath) > 0 {
		appIcon = iconPath[0]
	}
	err := beeep.Alert(config.Title, config.Message, appIcon)
	if err != nil {
		panic(err)
	}
}
