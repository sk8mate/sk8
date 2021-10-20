package domain

import (
	"fmt"

	"gorm.io/gorm"
	"sk8.town/backside/logger"
)

func dropTable(db *gorm.DB, data interface{}){
	if db.Migrator().HasTable(data) {
		err := db.Migrator().DropTable(data)
		if err != nil {
			logger.Error(err.Error())
			panic(err)
		}
		table := getTableName(db, data)
		logger.Info(fmt.Sprintf("Drop table %s ✓", table))
	}
}

func dropTables(db *gorm.DB) {
	dropTable(db, Spot{})
	dropTable(db, Filter{})
	dropTable(db, FilterValue{})
}

func migrateTable(db *gorm.DB, data interface{}){
	err := db.AutoMigrate(data)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	table := getTableName(db, data)
	logger.Info(fmt.Sprintf("Migrate %s ✓", table))
}
func autoMigrate(db *gorm.DB) {
	migrateTable(db, Spot{})
	migrateTable(db, Filter{})
	migrateTable(db, FilterValue{})
}

// TODO: move to utils
func getTableName(db *gorm.DB, Struct interface{}) string {
	stmt := &gorm.Statement{DB: db}
	if err := stmt.Parse(Struct); err != nil {
		logger.Error(err.Error())
		return "ERROR"
	} else {
		return stmt.Schema.Table
	}
}
