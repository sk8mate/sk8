package dto

import "sk8.town/backside/errs"

func (x *SpotRequest) Validate() *errs.AppError {
	if x.Coordinates == nil {
		return errs.NewValidationError(`Coordinates not provided`)
	}
	return nil
}
