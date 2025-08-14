package validation

import (
	"go.uber.org/dig"

	flam "github.com/happyhippyhippo/flam"
	config "github.com/happyhippyhippo/flam-config"
)

type provider struct{}

func NewProvider() flam.Provider {
	return &provider{}
}

func (provider) Id() string {
	return providerId
}

func (provider) Register(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	var e error
	provide := func(constructor any, opts ...dig.ProvideOption) bool {
		e = container.Provide(constructor, opts...)
		return e == nil
	}

	_ = provide(newEnglishLanguageCreator, dig.Group(LanguageCreatorGroup)) &&
		provide(newLanguageFactory) &&
		provide(newDefaultParserCreator, dig.Group(ParserCreatorGroup)) &&
		provide(newParserFactory) &&
		provide(newDefaultValidatorFuncCreator, dig.Group(ValidatorFuncCreatorGroup)) &&
		provide(newValidatorFuncFactory) &&
		provide(newFacade)

	return e
}

func (provider) Boot(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		configFacade config.Facade,
	) error {
		DefaultAnnotation = configFacade.String(PathDefaultAnnotation, DefaultAnnotation)
		DefaultStatus = configFacade.Int(PathDefaultStatus, DefaultStatus)
		DefaultTranslator = configFacade.String(PathDefaultTranslator, DefaultTranslator)
		DefaultParser = configFacade.String(PathDefaultParser, DefaultParser)
		DefaultLanguage = configFacade.String(PathDefaultLanguage, DefaultLanguage)

		return nil
	})
}

func (provider) Close(
	container *dig.Container,
) error {
	if container == nil {
		return newErrNilReference("container")
	}

	return container.Invoke(func(
		validatorFuncFactory validatorFuncFactory,
		parserFactory parserFactory,
		languageFactory languageFactory,
	) error {
		if e := validatorFuncFactory.Close(); e != nil {
			return e
		}

		if e := parserFactory.Close(); e != nil {
			return e
		}

		if e := languageFactory.Close(); e != nil {
			return e
		}

		return nil
	})
}
