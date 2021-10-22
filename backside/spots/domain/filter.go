package domain

type Filter struct {
	ID           int `gorm:"primary_key;AUTO_INCREMENT"`
	Name         string
	Type         string
	FilterValues []FilterValue
}

type FilterValue struct {
	ID       int `gorm:"primary_key;AUTO_INCREMENT"`
	Value    string
	FilterID int
}
