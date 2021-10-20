package domain

import (
	"fmt"
	"gorm.io/gorm"
	"sk8.town/backside/logger"
)

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

func dropTables(db *gorm.DB) {
	dropTable(db, Spot{})
	dropTable(db, Filter{})
	dropTable(db, FilterValue{})
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

func autoMigrate(db *gorm.DB) {
	migrateTable(db, Spot{})
	migrateTable(db, Filter{})
	migrateTable(db, FilterValue{})
}

func initializeTables(db *gorm.DB){
	lightingFilter:= Filter{
		Name: "Lighting",
		Type: "checkbox",
	}
	err := db.Create(&lightingFilter).Error
	if err != nil {
		panic(err)
	}

	obstaclesFilter:= Filter{
		Name: "Obstacles",
		Type: "multi_select",
	}
	err = db.Create(&obstaclesFilter).Error
	if err != nil {
		panic(err)
	}
	db.Find(&obstaclesFilter)

	obstaclesFilterValue1 := FilterValue{
		Value:    "Stairs",
		FilterId: obstaclesFilter.Id,
	}
	err = db.Create(&obstaclesFilterValue1).Error
	if err != nil {
		panic(err)
	}

	obstaclesFilterValue2 := FilterValue{
		Value:    "Mini-ramp",
		FilterId: obstaclesFilter.Id,
	}
	err = db.Create(&obstaclesFilterValue2).Error
	if err != nil {
		panic(err)
	}

	singleFilter := Filter{
		Name: "Somefilter",
		Type: "single_select",
	}
	err = db.Create(&singleFilter).Error
	if err != nil {
		panic(err)
	}
	db.Find(&singleFilter)

	singleFilterValue1 := FilterValue{
		Value:    "Some_value",
		FilterId: singleFilter.Id,
	}
	err = db.Create(&singleFilterValue1).Error
	if err != nil {
		panic(err)
	}

	singleFilterValue2 := FilterValue{
		Value:    "Some_other_value",
		FilterId: singleFilter.Id,
	}
	err = db.Create(&singleFilterValue2).Error
	if err != nil {
		panic(err)
	}
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
