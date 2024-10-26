package pkg

import (
	"fmt"
	"reflect"
)

type ChangesetValidator struct {
	IsValid bool
	Errors  map[string]string
}

func NewChangesetValidator() *ChangesetValidator {
	return &ChangesetValidator{
		IsValid: true,
		Errors:  make(map[string]string),
	}
}

func (c *ChangesetValidator) CastToString(param map[string]any) *ChangesetValidator {
	for key, value := range param {
		if vt := reflect.ValueOf(value).Kind(); vt != reflect.String {
			c.Errors[key] = fmt.Sprintf("field '%s' is not string type", key)
		}
	}

	if len(c.Errors) > 0 {
		c.IsValid = false
	}

	return c
}

func (c *ChangesetValidator) ValidateRequired(param map[string]any, fields []string) *ChangesetValidator {
	for _, field := range fields {
		value, ok := param[field]
		if !ok {
			c.Errors[field] = fmt.Sprintf("Missing required field '%s'", field)
		}

		if value == "" {
			c.Errors[field] = fmt.Sprintf("Empty required field '%s'", field)
		}
	}

	if len(c.Errors) > 0 {
		c.IsValid = false
	}

	return c
}