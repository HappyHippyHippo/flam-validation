package validation

import (
	flam "github.com/happyhippyhippo/flam"
)

type ValidatorFuncCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (ValidatorFunc, error)
}
