package validation

import (
	flam "github.com/happyhippyhippo/flam"
)

type LanguageCreator interface {
	Accept(config flam.Bag) bool
	Create(config flam.Bag) (Language, error)
}
