package validation

import (
	validator "github.com/go-playground/validator/v10"

	envelope "github.com/happyhippyhippo/flam-envelope"
	translation "github.com/happyhippyhippo/flam-translation"
)

type Parser interface {
	Translator() translation.Translator
	AddError(e string, code int)
	Parse(value any, errs validator.ValidationErrors) (envelope.Envelope, bool)
}
