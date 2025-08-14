package validation

import (
	"github.com/go-playground/validator/v10"

	flam "github.com/happyhippyhippo/flam"
)

type defaultValidatorFuncCreator struct {
	parserFactory   parserFactory
	languageFactory languageFactory
}

func newDefaultValidatorFuncCreator(
	parserFactory parserFactory,
	languageFactory languageFactory,
) ValidatorFuncCreator {
	return &defaultValidatorFuncCreator{
		parserFactory:   parserFactory,
		languageFactory: languageFactory,
	}
}

func (defaultValidatorFuncCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == ValidatorFuncDriverDefault
}

func (creator defaultValidatorFuncCreator) Create(
	config flam.Bag,
) (ValidatorFunc, error) {
	validate := validator.New()

	parserId := config.String("parser", DefaultParser)
	parser, e := creator.parserFactory.Get(parserId)
	if e != nil {
		return nil, e
	}

	if language := config.String("language", DefaultLanguage); language != "" {
		register, e := creator.languageFactory.Get(language)
		if e != nil {
			return nil, e
		}

		if e := register.Register(validate, parser.Translator()); e != nil {
			return nil, e
		}
	}

	return newDefaultValidatorFunc(
		validate,
		parser), nil
}
