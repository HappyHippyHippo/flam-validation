package validation

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
)

type languageFactory flam.Factory[Language]

type languageFactoryArgs struct {
	dig.In

	Creators      []LanguageCreator `group:"flam.validation.languages.creator"`
	FactoryConfig flam.FactoryConfig
}

func newLanguageFactory(
	args languageFactoryArgs,
) (languageFactory, error) {
	var creators []flam.ResourceCreator[Language]
	for _, creator := range args.Creators {
		creators = append(creators, creator)
	}

	return flam.NewFactory(
		creators,
		PathLanguages,
		args.FactoryConfig,
		nil)
}
