package validation

import (
	"github.com/go-playground/validator/v10"
	"go.uber.org/multierr"
)

// Playground — валидатор go-playground
type Playground struct {
	Validator *validator.Validate
}

// NewPlayground — создаем новый экземпляр валидатора
func NewPlayground() Playground {
	p := Playground{Validator: validator.New()}
	return p
}

// Validate — реализуем интерфейс Validator
func (p Playground) Validate(i interface{}) error {
	err := p.Validator.Struct(i)
	if err == nil {
		return nil
	}

	if fieldErrs, ok := err.(validator.ValidationErrors); ok {
		var combinedErr error

		for _, fieldErr := range fieldErrs {
			parameter := fieldErr.Param()

			multierr.AppendInto(&combinedErr, Error{
				CheckName: fieldErr.Tag(),
				Parameter: parameter,
				ValueName: fieldErr.Field(),
				Value:     fieldErr.Value(),
				cause:     fieldErr,
			})
		}

		return combinedErr
	}

	return Error{cause: err}
}
