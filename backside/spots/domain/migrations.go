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

func autoMigrate(db *gorm.DB) {
	migrateTable(db, Spot{})
	migrateTable(db, Filter{})
	migrateTable(db, FilterValue{})

	if !hasInitialValues(db) {
		initializeTables(db)
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

func hasInitialValues(db *gorm.DB) bool {
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

func initializeTables(db *gorm.DB) {
	lightingFilter := &Filter{
		Name: "Lighting",
		Type: "checkbox",
	}
	lightingFilter = addFilter(db, lightingFilter)

	obstaclesFilter := &Filter{
		Name: "Obstacles",
		Type: "multi_select",
	}
	obstaclesFilter = addFilter(db, obstaclesFilter)

	obstaclesFilterValue1 := &FilterValue{
		Value:    "Stairs",
		FilterId: obstaclesFilter.Id,
	}
	addFilterValue(db, obstaclesFilterValue1)

	obstaclesFilterValue2 := &FilterValue{
		Value:    "Mini-ramp",
		FilterId: obstaclesFilter.Id,
	}
	addFilterValue(db, obstaclesFilterValue2)

	singleFilter := &Filter{
		Name: "Somefilter",
		Type: "single_select",
	}
	singleFilter = addFilter(db, singleFilter)

	singleFilterValue1 := &FilterValue{
		Value:    "Some_value",
		FilterId: singleFilter.Id,
	}
	addFilterValue(db, singleFilterValue1)

	singleFilterValue2 := &FilterValue{
		Value:    "Some_other_value",
		FilterId: singleFilter.Id,
	}
	addFilterValue(db, singleFilterValue2)
}

func addFilter(db *gorm.DB, filter *Filter) *Filter {
	err := db.Create(filter).Error
	if err != nil {
		panic(err)
	}
	db.Find(&filter)
	return filter
}

func addFilterValue(db *gorm.DB, filterValue *FilterValue) {
	err := db.Create(&filterValue).Error
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
