package validation

import (
	flam "github.com/happyhippyhippo/flam"
)

type englishLanguageCreator struct{}

func newEnglishLanguageCreator() LanguageCreator {
	return &englishLanguageCreator{}
}

func (englishLanguageCreator) Accept(
	config flam.Bag,
) bool {
	return config.String("driver") == LanguageDriverEnglish
}

func (englishLanguageCreator) Create(
	_ flam.Bag,
) (Language, error) {
	return newEnglishLanguage(), nil
}
