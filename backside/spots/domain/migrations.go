package domain

import (
	"fmt"

	"gorm.io/gorm"
	"sk8.town/backside/logger"
)

func dropTables(db *gorm.DB) {
	dropTable(db, Spot{})
	dropTable(db, Filter{})
	dropTable(db, FilterValue{})
}

func autoMigrate(db *gorm.DB) {
	migrateTable(db, Spot{})
	migrateTable(db, Filter{})
	migrateTable(db, FilterValue{})

	if !hasInitData(db) {
		initData(db)
	}
}

func migrateTable(db *gorm.DB, data interface{}) {
	err := db.AutoMigrate(data)
	if err != nil {
		logger.Error(err.Error())
		panic(err)
	}
	table := getTableName(db, data)
	logger.Info(fmt.Sprintf("Migrate %s ✓", table))
}

func dropTable(db *gorm.DB, data interface{}) {
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

func initData(db *gorm.DB) {
	data := []Filter{
		{Name: "lighting", Type: "checkbox"},
		{Name: "obstacles", Type: "multi_select", FilterValues: []FilterValue{
			{Value: "mini-ramp"},
			{Value: "stairs"},
		}},
	}

	err := db.Create(data).Error

	if err != nil {
		panic(err)
	}
}

func hasInitData(db *gorm.DB) bool {
	var filters []*Filter
	err := db.Find(&filters).Error
	if err != nil {
		panic(err)
	}

	var filterValues []*FilterValue
	err = db.Find(&filterValues).Error
	if err != nil {
		panic(err)
	}

	return len(filters) > 0 && len(filterValues) > 0
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
