package main

import (
	"github.com/mdpe-ir/md-reminder/internal/app"
	"github.com/mdpe-ir/md-reminder/internal/configuration"
)

func main() {
	db := configuration.NewDatabase()
	newApp := app.NewApp(db)
	newApp.Run()
	defer db.Close()
}
