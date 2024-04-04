package validatorutil

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

type Validator struct {
}

func (v Validator) IsDate(fl validator.FieldLevel) bool {

	datePattern := `^\d{4}-\d{2}-\d{2}$`
	regex := regexp.MustCompile(datePattern)

	inter := fl.Field()
	slice, ok := inter.Interface().([]string)

	if !ok {
		dateString := fl.Field().String()
		return regex.MatchString(dateString)
	}
	for _, v := range slice {
		r := regex.MatchString(v)
		if !r {
			return r
		}
	}
	return true
}

func NewValidator() *validator.Validate {
	v := Validator{}
	validate := validator.New(validator.WithRequiredStructEnabled())
	_ = validate.RegisterValidation("IsDate", v.IsDate)
	return validate
}
