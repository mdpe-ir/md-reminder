package app

import (
	"fmt"
	"github.com/mdpe-ir/md-reminder/internal/constants"
	"github.com/mdpe-ir/md-reminder/internal/logic/deadline_logic"
	"github.com/mdpe-ir/md-reminder/internal/notification"
	"github.com/mdpe-ir/md-reminder/internal/systray"
	"github.com/mdpe-ir/md-reminder/internal/ui/pages"
	"github.com/mdpe-ir/md-reminder/internal/utils"
	"github.com/tidwall/buntdb"
	ptime "github.com/yaa110/go-persian-calendar"
	"io/ioutil"
	"log"
	"math"
	"time"
)

func NewApp(db *buntdb.DB) *App {
	return &App{DB: db}
}

type App struct {
	*buntdb.DB
}

func (app App) Run() {
	onExit := func() {
		// TODO: Log Exit
	}

	systray.Run(app.onReady, onExit)
}

func (app App) onReady() {
	var err error
	icon, err := ioutil.ReadFile("assets/md-reminder-system-tray-icon.png")
	if err != nil {
		log.Fatalf("Unable to read icon: %v", err)
	}
	systray.SetIcon(icon)
	systray.SetTitle(constants.AppName)
	systray.SetTooltip(constants.AppName)
	mShowNotify := systray.AddMenuItem(constants.Notify, constants.NotifyTooltip)
	mShowSetDeadline := systray.AddMenuItem(constants.SetDeadline, constants.SetDeadlineTooltip)
	mShowAbout := systray.AddMenuItem(constants.About, constants.AboutTooltip)
	mQuitOrig := systray.AddMenuItem(constants.Quit, constants.QuitTooltip)
	go func() {
		for {
			select {
			case <-mShowNotify.ClickedCh:
				deadlineTargetTime := deadline_logic.GetDeadlineDateTime(app.DB)
				daysCount := math.Round(deadlineTargetTime.Local().Sub(time.Now().Local()).Hours() / 24)
				persianTime := utils.PersianDateToString(ptime.New(deadlineTargetTime.Local()))
				notification.SendNotification(notification.Config{
					Title:   constants.NotificationTitle,
					Message: fmt.Sprintf(constants.NotificationBody, daysCount, persianTime),
				})
			case <-mQuitOrig.ClickedCh:
				fmt.Println("Requesting quit")
				systray.Quit()
				fmt.Println("Finished quitting")
				return
			case <-mShowSetDeadline.ClickedCh:
				pages.ShowSetDeadline(app.DB)
			case <-mShowAbout.ClickedCh:
				pages.ShowAbout()
			}

		}
	}()
}
