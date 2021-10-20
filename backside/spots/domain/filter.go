package domain

type Filter struct {
	Id   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Type string
}

type FilterValue struct {
	Id       int `gorm:"primary_key;AUTO_INCREMENT"`
	Value    string
	FilterId int    `gorm:"not null"`
	Filter   Filter `gorm:"foreignKey:FilterId"`
}
