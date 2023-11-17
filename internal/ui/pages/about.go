package pages

import (
	"github.com/mdpe-ir/md-reminder/internal/constants"
	"github.com/ncruces/zenity"
)

func ShowAbout() {
	aboutTxt := "MdReminder\n\n" +
		"Developed by: https://github.com/mdpe-ir\n\n" +
		"Source code: https://github.com/mdpe-ir/md-reminder\n\n" +
		"Version: " + constants.Version + "\n\n" +
		"Published: " + constants.PublishDate + "\n\n"

	zenity.Info(aboutTxt,
		zenity.Title("About"),
		zenity.InfoIcon)

}
