package validation

import (
	envelope "github.com/happyhippyhippo/flam-envelope"
)

type ValidatorFunc func(val any) (envelope.Envelope, bool)
