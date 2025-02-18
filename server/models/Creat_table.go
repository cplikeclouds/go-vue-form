package models

import (
	"fmt"
	"server/global"
)

func createTable(model interface{}) {
	db := global.DB
	if db.Migrator().HasTable(model) {
		return
	} else {
		err := db.AutoMigrate(model)
		if err != nil {
			panic(fmt.Sprintf("failed to migrate database: %v", err))
		}
	}
}
