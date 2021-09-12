package domain

import (
	"sk8.town/backside/errs"
)

//go:generate mockgen --build_flags=--mod=mod -destination=./mocks/spot_repository.go -package=mocks sk8.town/backside/spots SpotRepository
type SpotRepository interface {
	Add(Spot) (*Spot, *errs.AppError)
}

type SpotDb struct {
	client *sqlx.DB
}

func (repo SpotDb) Add(spot Spot) (*Spot, *errs.AppError) {

}

func NewSpotDb() SpotDb {
	return SpotDb{}
}

