package validate

import (
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
)

// KSUID does a simple regex string comparison for char type
// and length
func KSUID(fl validator.FieldLevel) bool {
	_, err := ksuid.Parse(fl.Field().String())
	return err == nil
}
