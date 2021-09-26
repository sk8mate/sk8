package domain

import (
	"fmt"

	"gorm.io/gorm"
	"sk8.town/backside/logger"
)

func dropTables(db *gorm.DB) {
	if db.Migrator().HasTable(Spot{}) {
		err := db.Migrator().DropTable(Spot{})
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		table := getTableName(db, Spot{})
		logger.Info(fmt.Sprintf("Drop table %s ✓", table))
	}
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(Spot{})
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	table := getTableName(db, Spot{})
	logger.Info(fmt.Sprintf("Migrate %s ✓", table))
}

// TODO: move to utils
func getTableName(db *gorm.DB, Struct interface{}) string {
	stmt := &gorm.Statement{DB: db}
	stmt.Parse(Struct)
	table := stmt.Schema.Table

	return table
}
