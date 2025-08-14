package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"

	translation "github.com/happyhippyhippo/flam-translation"
)

type englishLanguage struct{}

func newEnglishLanguage() Language {
	return &englishLanguage{}
}

func (englishLanguage) Register(
	validate *validator.Validate,
	translator translation.Translator,
) error {
	return en.RegisterDefaultTranslations(validate, translator)
}
