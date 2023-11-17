package utils

import (
	ptime "github.com/yaa110/go-persian-calendar"
	"strconv"
)

func PersianDateToString(pt ptime.Time) string {
	ptYear, ptMonth, ptDay := pt.Date()
	return strconv.Itoa(ptDay) + " " + ptMonth.String() + " " + strconv.Itoa(ptYear)
}
