package validate

import (
	"github.com/go-playground/validator/v10"

	modelv1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/model/v1"
	servicev1 "github.com/hampgoodwin/todo/gen/proto/go/to_do/service/v1"
)

// Validate wraps the logic for using go-playground/validator.
func Validate(i interface{}) error {
	v := validator.New()
	registerCustomValidations(v)
	if err := v.Struct(i); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		validationErrors := err.(validator.ValidationErrors)
		if len(validationErrors) > 0 {
			return err
		}
	}
	return nil
}

func Var(i interface{}, tag string) error {
	v := validator.New()
	registerCustomValidations(v)
	if err := v.Var(i, tag); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		validationErrors := err.(validator.ValidationErrors)
		if len(validationErrors) > 0 {
			return err
		}
	}
	return nil
}

func registerCustomValidations(v *validator.Validate) {
	// custom validations
	_ = v.RegisterValidation("KSUID", KSUID)

	// custom structs
	v.RegisterStructValidationMapRules(toDo, &modelv1.ToDo{})
	v.RegisterStructValidationMapRules(listToDoRequest, &servicev1.ListToDosRequest{})
}
