package app

import (
	"fmt"
	"github.com/mdpe-ir/md-reminder/internal/constants"
	"github.com/mdpe-ir/md-reminder/internal/systray"
	"io/ioutil"
	"log"
)

func Run() {
	onExit := func() {
		// TODO: Log Exit
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	var err error
	icon, err := ioutil.ReadFile("assets/md-reminder-system-tray-icon.png")
	if err != nil {
		log.Fatalf("Unable to read icon: %v", err)
	}
	systray.SetIcon(icon)
	systray.SetTitle(constants.AppName)
	systray.SetTooltip(constants.AppName)
	mShowNotify := systray.AddMenuItem(constants.Notify, constants.NotifyTooltip)
	mShowSettings := systray.AddMenuItem(constants.Settings, constants.SettingsTooltip)
	mShowAbout := systray.AddMenuItem(constants.About, constants.AboutTooltip)
	mQuitOrig := systray.AddMenuItem(constants.Quit, constants.QuitTooltip)
	go func() {
		for {
			select {
			case <-mShowNotify.ClickedCh:
				fmt.Println("Show Notify")
			case <-mQuitOrig.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				fmt.Println("Finished quitting")
				return
			case <-mShowSettings.ClickedCh:
				fmt.Println("Show Settings")

			case <-mShowAbout.ClickedCh:
				fmt.Println("Show About")
			}

		}
	}()
}
