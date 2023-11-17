package configuration

import (
	"github.com/tidwall/buntdb"
	"log"
)

func NewDatabase() *buntdb.DB {
	db, err := buntdb.Open("mdreminder.db")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
