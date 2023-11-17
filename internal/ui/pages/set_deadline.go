package pages

import (
	"fmt"
	"github.com/mdpe-ir/md-reminder/internal/constants"
	"github.com/mdpe-ir/md-reminder/internal/logic/deadline_logic"
	"github.com/mdpe-ir/md-reminder/internal/utils"
	"github.com/ncruces/zenity"
	"github.com/tidwall/buntdb"
	ptime "github.com/yaa110/go-persian-calendar"
)

func ShowSetDeadline(db *buntdb.DB) {

	defaultDeadlineTime := deadline_logic.GetDeadlineDateTime(db)

	calendar, err := zenity.Calendar(constants.SetDeadlineTitle,
		zenity.DefaultDate(defaultDeadlineTime.Year(), defaultDeadlineTime.Month(), defaultDeadlineTime.Day()),
		zenity.Title(constants.SetDeadline),
		zenity.OKLabel(constants.SetDeadlineOK),
		zenity.CancelLabel(constants.SetDeadlineCancel),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(constants.DBDeadlineKey, calendar.Local().String(), nil)
		return err
	})

	zenity.Info(fmt.Sprintf(constants.SetDeadlineInfoChange, utils.PersianDateToString(ptime.New(calendar))),
		zenity.Title(constants.Information),
		zenity.InfoIcon)

}
