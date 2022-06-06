package validation

import (
	"errors"
	"fmt"
)

// Error — ошибка валидации
type Error struct {
	CheckName   string
	Parameter   string
	ValueName   string
	Value       interface{}
	description string
	cause       error
}

// IsError — хелпер, проверяющий, что переданная ошибка является ошибкой валидации
func IsError(err error) bool {
	var validationErr Error
	return errors.As(err, &validationErr)
}

// Error — реализуем интерфейс builtin.error
func (err Error) Error() string {
	if err.CheckName != "" && err.ValueName != "" {
		if err.Parameter != "" {
			return fmt.Sprintf("'%s' failed the '%s=%s' check", err.ValueName, err.CheckName, err.Parameter)
		}

		return fmt.Sprintf("'%s' failed the '%s' check", err.ValueName, err.CheckName)
	}

	if err.CheckName != "" {
		return fmt.Sprintf("validation failed: %s", err.CheckName)
	}

	if err.cause != nil {
		return err.cause.Error()
	}

	return "unknown error"
}

func (err Error) WithCause(cause error) Error {
	err.cause = cause
	return err
}

func (err Error) WithValueName(name string) Error {
	err.ValueName = name
	return err
}

func (err Error) WithValue(value interface{}) Error {
	err.Value = value
	return err
}

func (err Error) WithNamedValue(name string, value interface{}) Error {
	err.ValueName = name
	err.Value = value
	return err
}

func (err Error) Cause() error {
	return err.cause
}

func (err Error) WithDescription(description string) Error {
	err.description = description
	return err
}

func (err Error) ValueAsString() string {
	switch v := err.Value.(type) {
	case string:
		return v
	case *string:
		if v != nil {
			return *v
		}
	}

	return ""
}
