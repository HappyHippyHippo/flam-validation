package validation

import (
	flam "github.com/happyhippyhippo/flam"
	translator "github.com/happyhippyhippo/flam-translation"
)

type defaultParserCreator struct {
	translatorFacade translator.Facade
}

func newDefaultParserCreator(
	translatorFacade translator.Facade,
) ParserCreator {
	return &defaultParserCreator{
		translatorFacade: translatorFacade,
	}
}

func (defaultParserCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == ParserDriverDefault
}

func (creator defaultParserCreator) Create(
	config flam.Bag,
) (Parser, error) {
	translatorId := config.String("translator", DefaultTranslator)
	translator, e := creator.translatorFacade.GetTranslator(translatorId)
	if e != nil {
		return nil, e
	}

	return newDefaultParser(
		translator,
		config.String("annotation", DefaultAnnotation),
		config.Int("status", DefaultStatus)), nil
}
