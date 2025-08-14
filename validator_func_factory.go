package validation

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type validatorFuncFactory flam.Factory[ValidatorFunc]

type validatorFuncFactoryArgs struct {
	dig.In

	Creators      []ValidatorFuncCreator `group:"flam.validation.validators.creator"`
	FactoryConfig flam.FactoryConfig
}

func newValidatorFuncFactory(
	args validatorFuncFactoryArgs,
) (validatorFuncFactory, error) {
	var creators []flam.ResourceCreator[ValidatorFunc]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathValidatorFuncs,
		args.FactoryConfig,
		nil)
}
