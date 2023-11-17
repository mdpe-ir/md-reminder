package deadline_logic

import (
	"github.com/mdpe-ir/md-reminder/internal/constants"
	"github.com/tidwall/buntdb"
	"strings"
	"time"
)

func GetDeadlineDateTime(db *buntdb.DB) time.Time {
	defaultDeadline := ""
	defaultDeadlineTime := time.Now().Local()

	db.View(func(tx *buntdb.Tx) error {
		defaultDeadline, _ = tx.Get(constants.DBDeadlineKey)
		return nil
	})

	if len(strings.TrimSpace(defaultDeadline)) > 0 {
		defaultDeadlineTime, _ = time.Parse("2006-01-02 15:04:05 -0700 -0700", strings.TrimSpace(defaultDeadline))
	}

	return defaultDeadlineTime
}
