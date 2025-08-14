package validation

type Facade interface {
	HasLanguage(id string) bool
	ListLanguages() []string
	GetLanguage(id string) (Language, error)
	AddLanguage(id string, language Language) error

	HasParser(id string) bool
	ListParsers() []string
	GetParser(id string) (Parser, error)
	AddParser(id string, parser Parser) error

	HasValidator(id string) bool
	ListValidators() []string
	GetValidator(id string) (ValidatorFunc, error)
	AddValidator(id string, validator ValidatorFunc) error
}

type validationFacade struct {
	languageFactory      languageFactory
	parserFactory        parserFactory
	validatorFuncFactory validatorFuncFactory
}

func newFacade(
	languageFactory languageFactory,
	parserFactory parserFactory,
	validatorFuncFactory validatorFuncFactory,
) Facade {
	return &validationFacade{
		languageFactory:      languageFactory,
		parserFactory:        parserFactory,
		validatorFuncFactory: validatorFuncFactory,
	}
}

func (facade validationFacade) HasLanguage(
	id string,
) bool {
	return facade.languageFactory.Has(id)
}

func (facade validationFacade) ListLanguages() []string {
	return facade.languageFactory.List()
}

func (facade validationFacade) GetLanguage(
	id string,
) (Language, error) {
	return facade.languageFactory.Get(id)
}

func (facade validationFacade) AddLanguage(
	id string,
	language Language,
) error {
	return facade.languageFactory.Add(id, language)
}

func (facade validationFacade) HasParser(
	id string,
) bool {
	return facade.parserFactory.Has(id)
}

func (facade validationFacade) ListParsers() []string {
	return facade.parserFactory.List()
}

func (facade validationFacade) GetParser(
	id string,
) (Parser, error) {
	return facade.parserFactory.Get(id)
}

func (facade validationFacade) AddParser(
	id string,
	parser Parser,
) error {
	return facade.parserFactory.Add(id, parser)
}

func (facade validationFacade) HasValidator(
	id string,
) bool {
	return facade.validatorFuncFactory.Has(id)
}

func (facade validationFacade) ListValidators() []string {
	return facade.validatorFuncFactory.List()
}

func (facade validationFacade) GetValidator(
	id string,
) (ValidatorFunc, error) {
	return facade.validatorFuncFactory.Get(id)
}

func (facade validationFacade) AddValidator(
	id string,
	validator ValidatorFunc,
) error {
	return facade.validatorFuncFactory.Add(id, validator)
}
