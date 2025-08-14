package validation

import (
	validator "github.com/go-playground/validator/v10"

	translation "github.com/happyhippyhippo/flam-translation"
)

type Language interface {
	Register(*validator.Validate, translation.Translator) error
}
