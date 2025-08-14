package validation

import (
	"github.com/go-playground/validator/v10"

	envelope "github.com/happyhippyhippo/flam-envelope"
)

func newDefaultValidatorFunc(
	validate *validator.Validate,
	parser Parser,
) ValidatorFunc {
	return func(value any) (envelope.Envelope, bool) {
		if errs := validate.Struct(value); errs != nil {
			return parser.Parse(value, errs.(validator.ValidationErrors))
		}
		return envelope.NewEnvelope(), true
	}
}
