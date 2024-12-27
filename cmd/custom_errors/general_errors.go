package custom_errors

import (
	"errors"
	"fmt"
)

type ConfigErrors struct {
	error
}

func (cErrors *ConfigErrors) Generate(message string, code int) *ConfigErrors {
	return &ConfigErrors{
		error: errors.New(fmt.Sprintf("%d : %s", code, message)),
	}
}
