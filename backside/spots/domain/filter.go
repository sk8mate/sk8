package domain

type Filter struct {
	ID   int `gorm:"primary_key;AUTO_INCREMENT"`
	Name string
	Type string
}

type FilterValue struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Value    string
	Filter   Filter `gorm:"foreignKey:FilterID;references:ID"`
	FilterID int    `gorm:"not null"`
}

type FilterWithFilterValue struct {
	FilterID    Filter
	FilterValue FilterValue
}
