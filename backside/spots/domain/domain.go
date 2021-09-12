package domain

type Spot struct {
	Name string `db:"name"`
	Address string `db:"address"`
	Coordinates struct {
		Lat  float64
		Long float64
	} `db:"coordinates"`
	Lighting bool `db:"lighting"`
	Friendly bool `db:"friendly"`
	Verified bool `db:"verified"`
}
