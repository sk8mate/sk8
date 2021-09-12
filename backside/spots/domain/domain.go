package domain

type Spot struct {
	Name        string      `db:"name"`
	Address     string      `db:"address"`
	Coordinates Coordinates `db:"coordinates"`
	Lighting    bool        `db:"lighting"`
	Friendly    bool        `db:"friendly"`
	Verified    bool        `db:"verified"`
}

type Coordinates struct {
	Lat  float64 `db:"lat"`
	Long float64 `db:"long"`
}
